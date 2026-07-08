import json
from pathlib import Path

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import numpy as np

INPUT = Path('/project/harness_metrics_from_logs.json')
OUT1 = Path('/project/harness_severity_stacked_comparison.png')
OUT2 = Path('/project/harness_conditions_20_comparison.png')
OUT3 = Path('/project/harness_conditions_grouped_by_run.png')

with INPUT.open() as f:
    data = json.load(f)

runs = data['runs']
labels = [r['run_label'] for r in runs]
n_runs = len(runs)
run_alphas = np.linspace(0.4, 1.0, n_runs)

# ---------- Chart 1: FP/FN bars stacked by severity ----------
severities = ['HIGH', 'MEDIUM', 'LOW', 'UNRATED']
severity_colors = {
    'HIGH': '#d62828',
    'MEDIUM': '#f77f00',
    'LOW': '#2a9d8f',
    'UNRATED': '#6c757d',
}

categories = []
for label in labels:
    categories.extend([f'{label} FP', f'{label} FN'])

stack_data = {s: [] for s in severities}
for run in runs:
    for key in ('false_positives', 'false_negatives'):
        overall_val = run.get('overall', {}).get(key)
        severity_total = 0
        has_any_severity_value = False
        for s in severities:
            sv = run.get('severity_levels', {}).get(s, {}).get(key)
            if sv is not None:
                has_any_severity_value = True
                severity_total += int(sv)

        fallback_to_unrated = (not has_any_severity_value) and (overall_val is not None)

        for s in severities:
            v = run.get('severity_levels', {}).get(s, {}).get(key)
            if v is None:
                if fallback_to_unrated and s == 'UNRATED':
                    stack_data[s].append(int(overall_val))
                else:
                    stack_data[s].append(0)
            else:
                stack_data[s].append(int(v))

        if has_any_severity_value and overall_val is not None and severity_total < int(overall_val):
            stack_data['UNRATED'][-1] += int(overall_val) - severity_total

x = np.arange(len(categories))
width = 0.65

fig, ax = plt.subplots(figsize=(12, 6))
bottom = np.zeros(len(categories))
for s in severities:
    values = np.array(stack_data[s])
    ax.bar(x, values, width, bottom=bottom, label=s, color=severity_colors[s])
    bottom += values

ax.set_title('False Positives / False Negatives by Run (Stacked by Severity)')
ax.set_xlabel('Run / Metric')
ax.set_ylabel('Count')
ax.set_xticks(x)
ax.set_xticklabels(categories, rotation=12, ha='right')
ax.legend(title='Severity', ncols=4)
ax.grid(axis='y', alpha=0.3)
# Symlog keeps zero while making smaller bars visible next to very large ones.
ax.set_yscale('symlog', linthresh=1)

totals = np.zeros(len(categories), dtype=float)
for s in severities:
    totals += np.array(stack_data[s], dtype=float)
for i, total in enumerate(totals):
    ax.text(x[i], total + (0.6 if total > 0 else 0.2), f'{int(total)}', ha='center', va='bottom', fontsize=9)

ax.text(0.01, -0.18, 'Symlog y-scale used so both high and low values remain visible.',
        transform=ax.transAxes, fontsize=9)

fig.tight_layout()
fig.savefig(OUT1, dpi=170)
plt.close(fig)

# ---------- Chart 2: 20-condition diverging FP/FN in one graph ----------
all_tests = [f'todo-{i:02d}' for i in range(1, 21)]
run_maps = [{row['test_id']: row for row in run.get('per_test', [])} for run in runs]

fig, ax = plt.subplots(figsize=(18, 8))
x = np.arange(len(all_tests))
group_width = 0.8
width = group_width / max(1, n_runs)
offsets = (np.arange(n_runs) - (n_runs - 1) / 2.0) * width
run_edge_colors = ['#111111', '#1d3557', '#2b9348', '#7b2cbf', '#6d6875']

for run_idx in range(n_runs):
    run_label = labels[run_idx]
    run_map = run_maps[run_idx]
    xpos = x + offsets[run_idx]
    run_alpha = float(run_alphas[run_idx])

    fp_stack = {s: [] for s in severities}
    fn_stack = {s: [] for s in severities}

    for t in all_tests:
        row = run_map.get(t, {})
        sev = row.get('severity', {})
        sev_fp = sev.get('false_positives', {})
        sev_fn = sev.get('false_negatives', {})
        row_fp_total = row.get('false_positives')
        row_fn_total = row.get('false_negatives')
        for s in severities:
            fp_stack[s].append(int(sev_fp.get(s, 0) or 0))
            fn_stack[s].append(int(sev_fn.get(s, 0) or 0))

        fp_known = sum(fp_stack[s][-1] for s in severities)
        fn_known = sum(fn_stack[s][-1] for s in severities)
        if row_fp_total is not None and fp_known < int(row_fp_total):
            fp_stack['UNRATED'][-1] += int(row_fp_total) - fp_known
        if row_fn_total is not None and fn_known < int(row_fn_total):
            fn_stack['UNRATED'][-1] += int(row_fn_total) - fn_known

    fp_bottom = np.zeros(len(all_tests), dtype=float)
    fn_bottom = np.zeros(len(all_tests), dtype=float)

    for s in severities:
        fp_vals = np.array(fp_stack[s], dtype=float)
        fn_vals = np.array(fn_stack[s], dtype=float)

        ax.bar(
            xpos,
            fp_vals,
            width,
            bottom=fp_bottom,
            color=severity_colors[s],
            edgecolor=run_edge_colors[run_idx % len(run_edge_colors)],
            linewidth=0.8,
            alpha=run_alpha,
            label=s if run_idx == 0 else None,
        )
        ax.bar(
            xpos,
            -fn_vals,
            width,
            bottom=-fn_bottom,
            color=severity_colors[s],
            edgecolor=run_edge_colors[run_idx % len(run_edge_colors)],
            linewidth=0.8,
            hatch='//',
            alpha=run_alpha,
        )

        fp_bottom += fp_vals
        fn_bottom += fn_vals

    for i in range(len(all_tests)):
        if fp_bottom[i] > 0:
            ax.text(xpos[i], fp_bottom[i] + 0.08, f'{int(fp_bottom[i])}', ha='center', va='bottom', fontsize=7)
        if fn_bottom[i] > 0:
            ax.text(xpos[i], -(fn_bottom[i] + 0.08), f'{int(fn_bottom[i])}', ha='center', va='top', fontsize=7)
        if fp_bottom[i] == 0 or fn_bottom[i] == 0:
            ax.scatter(
                xpos[i],
                0,
                s=26,
                marker='o',
                facecolor='#2b9348',
                edgecolor='white',
                linewidth=0.7,
                zorder=5,
            )

ax.axhline(0, color='#333333', linewidth=1.1)
ax.set_ylim(-2, 10)
ax.set_ylabel('Count (FP up, FN down)')
ax.set_xlabel('Condition')
ax.set_xticks(x)
ax.set_xticklabels(all_tests, rotation=50, ha='right')
ax.set_title('Per-Condition FP/FN Comparison Across Runs (20 Conditions), Stacked by Severity')
ax.grid(axis='y', alpha=0.3)

sev_handles, sev_labels = ax.get_legend_handles_labels()
legend1 = ax.legend(sev_handles, sev_labels, title='Severity', ncols=4, loc='upper left')
ax.add_artist(legend1)

run_labels = [f'{labels[i]} (series {i + 1})' for i in range(n_runs)]
run_handles = [
    plt.Line2D(
        [0],
        [0],
        color='white',
        marker='s',
        markerfacecolor='white',
        markeredgecolor=run_edge_colors[i % len(run_edge_colors)],
        alpha=float(run_alphas[i]),
        markersize=10,
        linewidth=0,
    )
    for i in range(n_runs)
]
legend2 = ax.legend(run_handles, run_labels, title='Run', loc='upper right')
ax.add_artist(legend2)

ax.text(0.995, 0.92, 'Solid = FP, hatched = FN', transform=ax.transAxes, ha='right', va='top', fontsize=9)
fig.text(0.01, 0.01, 'All per-condition FP/FN values are populated from git-history artifacts.', fontsize=9)
fig.tight_layout(rect=[0, 0.03, 1, 0.98])
fig.savefig(OUT2, dpi=170)
plt.close(fig)

# ---------- Chart 3: 20 conditions grouped by run (small multiples) ----------
fig, axes = plt.subplots(n_runs, 1, figsize=(18, 3.2 * max(1, n_runs)), sharex=True)
if n_runs == 1:
    axes = [axes]

for run_idx in range(n_runs):
    ax = axes[run_idx]
    run_map = run_maps[run_idx]
    label = labels[run_idx]

    fp_vals = []
    fn_vals = []
    for t in all_tests:
        row = run_map.get(t, {})
        sev = row.get('severity', {})

        # Prefer explicit per-test totals, fallback to severity sums.
        fp = row.get('false_positives')
        if fp is None:
            fp = sum(int((sev.get('false_positives', {}).get(s) or 0)) for s in severities)

        fn = row.get('false_negatives')
        if fn is None:
            fn = sum(int((sev.get('false_negatives', {}).get(s) or 0)) for s in severities)

        fp_vals.append(int(fp or 0))
        fn_vals.append(int(fn or 0))

    x = np.arange(len(all_tests))
    ax.bar(x, fp_vals, color='#2a9d8f', alpha=0.85, label='FP')
    ax.bar(x, -np.array(fn_vals), color='#d62828', alpha=0.85, label='FN')
    ax.axhline(0, color='#333333', linewidth=1.0)
    ax.grid(axis='y', alpha=0.25)
    ax.set_ylabel('Count')
    ax.set_title(label)

    for i, v in enumerate(fp_vals):
        if v > 0:
            ax.text(i, v + 0.08, str(v), ha='center', va='bottom', fontsize=7)
    for i, v in enumerate(fn_vals):
        if v > 0:
            ax.text(i, -(v + 0.08), str(v), ha='center', va='top', fontsize=7)

axes[-1].set_xticks(np.arange(len(all_tests)))
axes[-1].set_xticklabels(all_tests, rotation=50, ha='right')
axes[0].legend(loc='upper right')
fig.suptitle('20 Conditions Grouped by Run (FP up, FN down)', y=0.995)
fig.tight_layout(rect=[0, 0.02, 1, 0.98])
fig.savefig(OUT3, dpi=170)
plt.close(fig)

print(f'Saved {OUT1}')
print(f'Saved {OUT2}')
print(f'Saved {OUT3}')
