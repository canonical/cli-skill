"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
exports.findTranscriptDirs = findTranscriptDirs;
exports.parseTranscript = parseTranscript;
exports.getAllSessions = getAllSessions;
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
/**
 * Find all workspace storage directories that contain Copilot chat transcripts.
 */
function findTranscriptDirs(vscodeUserDataPath) {
    const workspaceStoragePath = path.join(vscodeUserDataPath, 'workspaceStorage');
    const dirs = [];
    if (!fs.existsSync(workspaceStoragePath)) {
        return dirs;
    }
    for (const entry of fs.readdirSync(workspaceStoragePath)) {
        const transcriptsDir = path.join(workspaceStoragePath, entry, 'GitHub.copilot-chat', 'transcripts');
        if (fs.existsSync(transcriptsDir)) {
            dirs.push(transcriptsDir);
        }
    }
    return dirs;
}
/**
 * Parse a single JSONL transcript file into structured session data.
 */
function parseTranscript(filePath) {
    let content;
    try {
        content = fs.readFileSync(filePath, 'utf-8');
    }
    catch {
        return null;
    }
    const lines = content.split('\n').filter(l => l.trim());
    if (lines.length === 0) {
        return null;
    }
    let sessionId = path.basename(filePath, '.jsonl');
    let startTime = '';
    const prompts = [];
    let firstAssistantResponse = '';
    const subAgents = [];
    for (const line of lines) {
        let entry;
        try {
            entry = JSON.parse(line);
        }
        catch {
            continue;
        }
        if (entry.type === 'session.start' && entry.data) {
            sessionId = entry.data.sessionId || sessionId;
            startTime = entry.data.startTime || entry.timestamp || '';
        }
        if (entry.type === 'user.message' && entry.data) {
            const rawContent = entry.data.content || '';
            const prompt = extractUserRequest(rawContent);
            if (prompt && prompt.trim()) {
                prompts.push(prompt.trim());
            }
        }
        if (entry.type === 'assistant.message' && entry.data && !firstAssistantResponse) {
            const resp = entry.data.content || '';
            if (resp.trim()) {
                firstAssistantResponse = resp.trim();
            }
        }
        if (entry.type === 'tool.execution_start' && entry.data) {
            const toolName = entry.data.toolName;
            if (toolName === 'runSubagent' || toolName === 'explore_subagent') {
                let args = {};
                try {
                    const rawArgs = entry.data.arguments;
                    args = typeof rawArgs === 'string' ? JSON.parse(rawArgs) : rawArgs || {};
                }
                catch { /* ignore */ }
                subAgents.push({
                    description: args.description || '',
                    agentName: args.agentName || 'default',
                    promptPreview: (args.prompt || args.query || '').slice(0, 150),
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
function extractUserRequest(content) {
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
function deriveTitle(firstPrompt) {
    const cleaned = firstPrompt.replace(/\n/g, ' ').trim();
    if (cleaned.length <= 80) {
        return cleaned;
    }
    return cleaned.slice(0, 77) + '...';
}
/**
 * Get all parsed sessions from all available transcript directories.
 */
function getAllSessions(vscodeUserDataPath) {
    const transcriptDirs = findTranscriptDirs(vscodeUserDataPath);
    const sessions = [];
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
        if (!a.startTime)
            return 1;
        if (!b.startTime)
            return -1;
        return new Date(b.startTime).getTime() - new Date(a.startTime).getTime();
    });
    return sessions;
}
//# sourceMappingURL=parser.js.map