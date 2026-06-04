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
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = __importStar(require("vscode"));
const path = __importStar(require("path"));
const parser_1 = require("./parser");
const htmlExport_1 = require("./htmlExport");
function activate(context) {
    context.subscriptions.push(vscode.commands.registerCommand('copilotExport.exportSessions', () => exportSessions(context)), vscode.commands.registerCommand('copilotExport.exportSessionsHtml', () => exportSessionsHtml(context)));
}
async function exportSessions(context) {
    // VS Code user data path is typically the parent of globalStoragePath
    const globalStoragePath = context.globalStorageUri.fsPath;
    const userDataPath = path.resolve(globalStoragePath, '..', '..');
    const sessions = (0, parser_1.getAllSessions)(userDataPath);
    if (sessions.length === 0) {
        vscode.window.showInformationMessage('No Copilot chat sessions found.');
        return;
    }
    const markdown = generateMarkdown(sessions);
    // Ask where to save
    const defaultUri = vscode.workspace.workspaceFolders?.[0]?.uri;
    const saveUri = await vscode.window.showSaveDialog({
        defaultUri: defaultUri
            ? vscode.Uri.joinPath(defaultUri, 'copilot-sessions.md')
            : undefined,
        filters: { Markdown: ['md'] },
        title: 'Save Copilot Sessions Export',
    });
    if (!saveUri) {
        return;
    }
    await vscode.workspace.fs.writeFile(saveUri, Buffer.from(markdown, 'utf-8'));
    vscode.window.showInformationMessage(`Exported ${sessions.length} sessions to ${path.basename(saveUri.fsPath)}`);
    // Open the file
    const doc = await vscode.workspace.openTextDocument(saveUri);
    await vscode.window.showTextDocument(doc);
}
async function exportSessionsHtml(context) {
    const globalStoragePath = context.globalStorageUri.fsPath;
    const userDataPath = path.resolve(globalStoragePath, '..', '..');
    const sessions = (0, parser_1.getAllSessions)(userDataPath);
    if (sessions.length === 0) {
        vscode.window.showInformationMessage('No Copilot chat sessions found.');
        return;
    }
    const html = (0, htmlExport_1.generateHtml)(sessions);
    const defaultUri = vscode.workspace.workspaceFolders?.[0]?.uri;
    const saveUri = await vscode.window.showSaveDialog({
        defaultUri: defaultUri
            ? vscode.Uri.joinPath(defaultUri, 'copilot-sessions.html')
            : undefined,
        filters: { HTML: ['html'] },
        title: 'Save Copilot Sessions HTML Export',
    });
    if (!saveUri) {
        return;
    }
    await vscode.workspace.fs.writeFile(saveUri, Buffer.from(html, 'utf-8'));
    vscode.window.showInformationMessage(`Exported ${sessions.length} sessions to ${path.basename(saveUri.fsPath)}`);
    const doc = await vscode.workspace.openTextDocument(saveUri);
    await vscode.window.showTextDocument(doc);
}
function generateMarkdown(sessions) {
    const lines = [];
    lines.push('# Copilot Chat Sessions Export');
    lines.push('');
    lines.push(`> Generated on ${new Date().toISOString().split('T')[0]}`);
    lines.push(`> Total sessions: ${sessions.length}`);
    lines.push('');
    lines.push('---');
    lines.push('');
    // Table of contents
    lines.push('## Table of Contents');
    lines.push('');
    for (let i = 0; i < sessions.length; i++) {
        const s = sessions[i];
        const date = s.startTime ? new Date(s.startTime).toLocaleDateString() : 'Unknown date';
        const anchor = `session-${i + 1}`;
        lines.push(`${i + 1}. [${s.title}](#${anchor}) — *${date}*`);
    }
    lines.push('');
    lines.push('---');
    lines.push('');
    // Each session
    for (let i = 0; i < sessions.length; i++) {
        const s = sessions[i];
        const date = s.startTime
            ? new Date(s.startTime).toLocaleString()
            : 'Unknown';
        lines.push(`## Session ${i + 1}`);
        lines.push(`### ${s.title}`);
        lines.push('');
        lines.push(`**Date:** ${date}  `);
        lines.push(`**Session ID:** \`${s.sessionId}\`  `);
        lines.push(`**Prompts:** ${s.prompts.length}`);
        lines.push('');
        // Summary from first assistant response
        if (s.firstAssistantResponse) {
            lines.push('**Summary:**');
            const summary = s.firstAssistantResponse.slice(0, 300);
            lines.push(summary + (s.firstAssistantResponse.length > 300 ? '...' : ''));
            lines.push('');
        }
        // List of prompts
        lines.push('**Prompts:**');
        lines.push('');
        for (let j = 0; j < s.prompts.length; j++) {
            lines.push(`**${j + 1}.**`);
            lines.push('```');
            lines.push(truncateToWords(s.prompts[j], 300));
            lines.push('```');
            lines.push('');
        }
        lines.push('---');
        lines.push('');
    }
    return lines.join('\n');
}
function deactivate() { }
function truncateToWords(text, maxWords) {
    const words = text.split(/\s+/);
    if (words.length <= maxWords) {
        return text;
    }
    return words.slice(0, maxWords).join(' ') + ' [...]';
}
//# sourceMappingURL=extension.js.map