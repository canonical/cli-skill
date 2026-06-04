import * as fs from 'fs';
import * as path from 'path';

export interface SubAgentCall {
  description: string;
  agentName: string;
  promptPreview: string;
  timestamp: string;
}

export interface SessionData {
  sessionId: string;
  startTime: string;
  title: string;
  prompts: string[];
  firstAssistantResponse: string;
  subAgents: SubAgentCall[];
}

interface TranscriptEntry {
  type: string;
  id: string;
  timestamp: string;
  parentId: string | null;
  data?: Record<string, unknown>;
}

/**
 * Find all workspace storage directories that contain Copilot chat transcripts.
 */
export function findTranscriptDirs(vscodeUserDataPath: string): string[] {
  const workspaceStoragePath = path.join(vscodeUserDataPath, 'workspaceStorage');
  const dirs: string[] = [];

  if (!fs.existsSync(workspaceStoragePath)) {
    return dirs;
  }

  for (const entry of fs.readdirSync(workspaceStoragePath)) {
    const transcriptsDir = path.join(
      workspaceStoragePath,
      entry,
      'GitHub.copilot-chat',
      'transcripts'
    );
    if (fs.existsSync(transcriptsDir)) {
      dirs.push(transcriptsDir);
    }
  }

  return dirs;
}

/**
 * Parse a single JSONL transcript file into structured session data.
 */
export function parseTranscript(filePath: string): SessionData | null {
  let content: string;
  try {
    content = fs.readFileSync(filePath, 'utf-8');
  } catch {
    return null;
  }

  const lines = content.split('\n').filter(l => l.trim());
  if (lines.length === 0) {
    return null;
  }

  let sessionId = path.basename(filePath, '.jsonl');
  let startTime = '';
  const prompts: string[] = [];
  let firstAssistantResponse = '';
  const subAgents: SubAgentCall[] = [];

  for (const line of lines) {
    let entry: TranscriptEntry;
    try {
      entry = JSON.parse(line);
    } catch {
      continue;
    }

    if (entry.type === 'session.start' && entry.data) {
      sessionId = (entry.data.sessionId as string) || sessionId;
      startTime = (entry.data.startTime as string) || entry.timestamp || '';
    }

    if (entry.type === 'user.message' && entry.data) {
      const rawContent = (entry.data.content as string) || '';
      const prompt = extractUserRequest(rawContent);
      if (prompt && prompt.trim()) {
        prompts.push(prompt.trim());
      }
    }

    if (entry.type === 'assistant.message' && entry.data && !firstAssistantResponse) {
      const resp = (entry.data.content as string) || '';
      if (resp.trim()) {
        firstAssistantResponse = resp.trim();
      }
    }

    if (entry.type === 'tool.execution_start' && entry.data) {
      const toolName = entry.data.toolName as string;
      if (toolName === 'runSubagent' || toolName === 'explore_subagent') {
        let args: Record<string, unknown> = {};
        try {
          const rawArgs = entry.data.arguments;
          args = typeof rawArgs === 'string' ? JSON.parse(rawArgs) : (rawArgs as Record<string, unknown>) || {};
        } catch { /* ignore */ }
        subAgents.push({
          description: (args.description as string) || '',
          agentName: (args.agentName as string) || 'default',
          promptPreview: ((args.prompt as string) || (args.query as string) || '').slice(0, 150),
          timestamp: entry.timestamp || '',
        });
      }
    }
  }

  if (prompts.length === 0) {
    return null;
  }

  // Derive title from first prompt (truncated)
  const title = deriveTitle(prompts[0]);

  return { sessionId, startTime, title, prompts, firstAssistantResponse, subAgents };
}

/**
 * Extract the user's actual request from potentially XML-wrapped content.
 */
function extractUserRequest(content: string): string {
  // If content has <userRequest> tags, extract just that
  const userReqMatch = content.match(/<userRequest>([\s\S]*?)<\/userRequest>/);
  if (userReqMatch) {
    return userReqMatch[1].trim();
  }

  // Strip any XML-like wrapper tags that Copilot adds
  const stripped = content
    .replace(/<context>[\s\S]*?<\/context>/g, '')
    .replace(/<editorContext>[\s\S]*?<\/editorContext>/g, '')
    .replace(/<reminderInstructions>[\s\S]*?<\/reminderInstructions>/g, '')
    .replace(/<environment_info>[\s\S]*?<\/environment_info>/g, '')
    .replace(/<workspace_info>[\s\S]*?<\/workspace_info>/g, '')
    .replace(/<userMemory>[\s\S]*?<\/userMemory>/g, '')
    .replace(/<sessionMemory>[\s\S]*?<\/sessionMemory>/g, '')
    .replace(/<repoMemory>[\s\S]*?<\/repoMemory>/g, '')
    .replace(/<availableDeferredTools>[\s\S]*?<\/availableDeferredTools>/g, '')
    .replace(/<instructions>[\s\S]*?<\/instructions>/g, '')
    .replace(/<securityRequirements>[\s\S]*?<\/securityRequirements>/g, '')
    .replace(/<[a-zA-Z_]+>[\s\S]*?<\/[a-zA-Z_]+>/g, '')
    .trim();

  return stripped || content;
}

/**
 * Derive a short title from the first user prompt.
 */
function deriveTitle(firstPrompt: string): string {
  const cleaned = firstPrompt.replace(/\n/g, ' ').trim();
  if (cleaned.length <= 80) {
    return cleaned;
  }
  return cleaned.slice(0, 77) + '...';
}

/**
 * Get all parsed sessions from all available transcript directories.
 */
export function getAllSessions(vscodeUserDataPath: string): SessionData[] {
  const transcriptDirs = findTranscriptDirs(vscodeUserDataPath);
  const sessions: SessionData[] = [];

  for (const dir of transcriptDirs) {
    const files = fs.readdirSync(dir).filter(f => f.endsWith('.jsonl'));
    for (const file of files) {
      const session = parseTranscript(path.join(dir, file));
      if (session) {
        sessions.push(session);
      }
    }
  }

  // Sort by start time (newest first)
  sessions.sort((a, b) => {
    if (!a.startTime) return 1;
    if (!b.startTime) return -1;
    return new Date(b.startTime).getTime() - new Date(a.startTime).getTime();
  });

  return sessions;
}
