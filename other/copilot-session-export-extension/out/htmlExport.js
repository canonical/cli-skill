"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateHtml = generateHtml;
function escapeHtml(text) {
    return text
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;');
}
function truncateToWords(text, maxWords) {
    const words = text.split(/\s+/);
    if (words.length <= maxWords) {
        return text;
    }
    return words.slice(0, maxWords).join(' ') + ' [...]';
}
function generateSubAgentSvg(session) {
    if (session.subAgents.length === 0) {
        return '';
    }
    const agents = session.subAgents;
    const nodeHeight = 44;
    const nodeWidth = 220;
    const gapY = 20;
    const startY = 70;
    const centerX = 300;
    // Group agents that were launched together (within 2s of each other)
    const groups = [];
    let currentGroup = [agents[0]];
    for (let i = 1; i < agents.length; i++) {
        const prev = new Date(agents[i - 1].timestamp).getTime();
        const curr = new Date(agents[i].timestamp).getTime();
        if (curr - prev < 2000) {
            currentGroup.push(agents[i]);
        }
        else {
            groups.push(currentGroup);
            currentGroup = [agents[i]];
        }
    }
    groups.push(currentGroup);
    const totalRows = groups.reduce((sum, g) => sum + Math.max(g.length, 1), 0) + groups.length;
    const svgHeight = startY + totalRows * (nodeHeight + gapY) + 40;
    const svgWidth = Math.max(600, agents.length > 2 ? 700 : 600);
    let svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${svgWidth} ${svgHeight}" class="subagent-svg">`;
    svg += `<defs>
    <marker id="arrowhead" markerWidth="10" markerHeight="7" refX="10" refY="3.5" orient="auto">
      <polygon points="0 0, 10 3.5, 0 7" fill="#6b7280"/>
    </marker>
  </defs>`;
    // Main session node
    svg += `<rect x="${centerX - 80}" y="10" width="160" height="40" rx="8" fill="#2563eb" />`;
    svg += `<text x="${centerX}" y="35" text-anchor="middle" fill="white" font-size="13" font-family="Ubuntu Sans, sans-serif">Main Session</text>`;
    let yOffset = startY;
    for (const group of groups) {
        const isParallel = group.length > 1;
        const groupWidth = group.length * (nodeWidth + 30) - 30;
        const groupStartX = centerX - groupWidth / 2;
        if (isParallel) {
            // Fan-out label
            svg += `<text x="${centerX}" y="${yOffset - 5}" text-anchor="middle" fill="#6b7280" font-size="11" font-family="Ubuntu Sans, sans-serif" font-style="italic">parallel</text>`;
        }
        for (let i = 0; i < group.length; i++) {
            const agent = group[i];
            const nodeX = isParallel ? groupStartX + i * (nodeWidth + 30) : centerX - nodeWidth / 2;
            const nodeY = yOffset;
            // Arrow from main to node
            const arrowStartY = yOffset > startY ? yOffset - gapY : 50;
            svg += `<line x1="${centerX}" y1="${arrowStartY}" x2="${nodeX + nodeWidth / 2}" y2="${nodeY}" stroke="#6b7280" stroke-width="1.5" marker-end="url(#arrowhead)" />`;
            // Agent node
            const color = agent.agentName === 'default' ? '#7c3aed' : '#059669';
            svg += `<rect x="${nodeX}" y="${nodeY}" width="${nodeWidth}" height="${nodeHeight}" rx="6" fill="${color}" opacity="0.9" />`;
            const label = escapeHtml(agent.description || agent.agentName).slice(0, 28);
            svg += `<text x="${nodeX + nodeWidth / 2}" y="${nodeY + 18}" text-anchor="middle" fill="white" font-size="11" font-family="Ubuntu Sans, sans-serif" font-weight="bold">${label}</text>`;
            const detail = escapeHtml(agent.promptPreview).slice(0, 35);
            svg += `<text x="${nodeX + nodeWidth / 2}" y="${nodeY + 34}" text-anchor="middle" fill="#e2e8f0" font-size="10" font-family="Ubuntu Mono, monospace">${detail}...</text>`;
        }
        yOffset += nodeHeight + gapY + (isParallel ? 20 : 0);
    }
    svg += '</svg>';
    return svg;
}
function generateHtml(sessions) {
    const date = new Date().toISOString().split('T')[0];
    let toc = '';
    let content = '';
    for (let i = 0; i < sessions.length; i++) {
        const s = sessions[i];
        const sessionDate = s.startTime
            ? new Date(s.startTime).toLocaleString()
            : 'Unknown';
        const shortDate = s.startTime
            ? new Date(s.startTime).toLocaleDateString()
            : '';
        toc += `<li><a href="#session-${i + 1}">${escapeHtml(s.title)}</a> <span class="date">${shortDate}</span></li>\n`;
        content += `<section id="session-${i + 1}" class="session">
  <h2>Session ${i + 1}</h2>
  <h3>${escapeHtml(s.title)}</h3>
  <div class="meta">
    <span><strong>Date:</strong> ${sessionDate}</span>
    <span><strong>Prompts:</strong> ${s.prompts.length}</span>
    <span><strong>Sub-agents:</strong> ${s.subAgents.length}</span>
  </div>
`;
        if (s.firstAssistantResponse) {
            const summary = escapeHtml(truncateToWords(s.firstAssistantResponse, 60));
            content += `  <div class="summary"><strong>Summary:</strong> ${summary}</div>\n`;
        }
        // Sub-agent SVG
        if (s.subAgents.length > 0) {
            content += `  <div class="subagent-flow">
    <h4>Sub-agent Flow</h4>
    ${generateSubAgentSvg(s)}
  </div>\n`;
        }
        // Prompts
        content += `  <div class="prompts">
    <h4>Prompts</h4>\n`;
        for (let j = 0; j < s.prompts.length; j++) {
            const prompt = escapeHtml(truncateToWords(s.prompts[j], 300));
            content += `    <div class="prompt">
      <span class="prompt-num">${j + 1}</span>
      <pre><code>${prompt}</code></pre>
    </div>\n`;
        }
        content += `  </div>\n</section>\n<hr/>\n`;
    }
    return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Copilot Sessions Export</title>
  <style>
    @import url('https://fonts.googleapis.com/css2?family=Ubuntu+Mono:wght@400;700&family=Ubuntu+Sans:wght@300;400;500;600;700&display=swap');

    :root {
      --bg: #f8fafc;
      --surface: #ffffff;
      --text: #1e293b;
      --text-muted: #64748b;
      --border: #e2e8f0;
      --accent: #2563eb;
      --accent-light: #dbeafe;
      --code-bg: #1e293b;
      --code-text: #e2e8f0;
    }

    * { margin: 0; padding: 0; box-sizing: border-box; }

    body {
      font-family: 'Ubuntu Sans', sans-serif;
      background: var(--bg);
      color: var(--text);
      line-height: 1.6;
      padding: 2rem;
      max-width: 960px;
      margin: 0 auto;
    }

    header {
      text-align: center;
      margin-bottom: 3rem;
      padding-bottom: 2rem;
      border-bottom: 2px solid var(--border);
    }

    header h1 {
      font-size: 2rem;
      font-weight: 700;
      margin-bottom: 0.5rem;
    }

    header .subtitle {
      color: var(--text-muted);
      font-size: 0.95rem;
    }

    nav {
      background: var(--surface);
      border: 1px solid var(--border);
      border-radius: 12px;
      padding: 1.5rem 2rem;
      margin-bottom: 3rem;
    }

    nav h2 {
      font-size: 1.1rem;
      margin-bottom: 0.75rem;
      color: var(--accent);
    }

    nav ol {
      padding-left: 1.5rem;
    }

    nav li {
      margin-bottom: 0.4rem;
      font-size: 0.9rem;
    }

    nav li a {
      color: var(--accent);
      text-decoration: none;
    }

    nav li a:hover { text-decoration: underline; }

    nav .date {
      color: var(--text-muted);
      font-size: 0.8rem;
      margin-left: 0.5rem;
    }

    .session {
      background: var(--surface);
      border: 1px solid var(--border);
      border-radius: 12px;
      padding: 2rem;
      margin-bottom: 1.5rem;
    }

    .session h2 {
      font-size: 0.85rem;
      text-transform: uppercase;
      letter-spacing: 0.05em;
      color: var(--text-muted);
      margin-bottom: 0.25rem;
    }

    .session h3 {
      font-size: 1.25rem;
      font-weight: 600;
      margin-bottom: 1rem;
    }

    .meta {
      display: flex;
      gap: 1.5rem;
      font-size: 0.85rem;
      color: var(--text-muted);
      margin-bottom: 1rem;
      flex-wrap: wrap;
    }

    .summary {
      background: var(--accent-light);
      border-radius: 8px;
      padding: 0.75rem 1rem;
      font-size: 0.9rem;
      margin-bottom: 1.5rem;
    }

    .subagent-flow {
      margin: 1.5rem 0;
    }

    .subagent-flow h4 {
      font-size: 0.9rem;
      color: var(--text-muted);
      margin-bottom: 0.75rem;
    }

    .subagent-svg {
      width: 100%;
      max-width: 700px;
      height: auto;
      display: block;
      margin: 0 auto;
    }

    .prompts h4 {
      font-size: 0.9rem;
      color: var(--text-muted);
      margin-bottom: 0.75rem;
    }

    .prompt {
      display: flex;
      align-items: flex-start;
      gap: 0.75rem;
      margin-bottom: 0.75rem;
    }

    .prompt-num {
      flex-shrink: 0;
      width: 28px;
      height: 28px;
      background: var(--accent);
      color: white;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 0.75rem;
      font-weight: 600;
      margin-top: 0.5rem;
    }

    .prompt pre {
      flex: 1;
      background: var(--code-bg);
      color: var(--code-text);
      border-radius: 8px;
      padding: 0.75rem 1rem;
      overflow-x: auto;
      font-size: 0.85rem;
      line-height: 1.5;
    }

    .prompt code {
      font-family: 'Ubuntu Mono', monospace;
      white-space: pre-wrap;
      word-break: break-word;
    }

    hr {
      border: none;
      border-top: 1px solid var(--border);
      margin: 2rem 0;
    }

    @media (max-width: 640px) {
      body { padding: 1rem; }
      .session { padding: 1.25rem; }
      .meta { flex-direction: column; gap: 0.5rem; }
    }
  </style>
</head>
<body>
  <header>
    <h1>Copilot Chat Sessions</h1>
    <p class="subtitle">Generated on ${date} &mdash; ${sessions.length} sessions</p>
  </header>

  <nav>
    <h2>Table of Contents</h2>
    <ol>
      ${toc}
    </ol>
  </nav>

  ${content}
</body>
</html>`;
}
//# sourceMappingURL=htmlExport.js.map