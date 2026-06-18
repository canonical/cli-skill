#!/usr/bin/env python3
"""
Analyze cli-review results from completed test variations.
Reads existing cli-review output and generates ANALYSIS.json
"""

import json
import re
from pathlib import Path
from dataclasses import dataclass, asdict
from typing import List, Dict, Optional

@dataclass
class TestResult:
    variation_id: str
    violations_injected: int
    findings_detected: int
    false_negatives: int
    severity_breakdown: Dict
    detected_issues: List[Dict]

def extract_findings_from_report(report_path: Path) -> List[Dict]:
    """Extract findings from cli-review.md report."""
    if not report_path.exists():
        return []
    
    findings = []
    try:
        content = report_path.read_text()
        
        # Extract severity summary from the report
        # Look for patterns like "High | 1 |" in the summary table
        severity_pattern = r'\|\s*(High|Medium|Low|Unrated)\s*\|\s*(\d+)\s*\|'
        matches = re.findall(severity_pattern, content)
        
        for severity, count in matches:
            findings.append({
                "severity": severity,
                "count": int(count)
            })
    except Exception as e:
        print(f"Error parsing {report_path}: {e}")
    
    return findings

def read_violation_metadata(metadata_path: Path) -> Dict:
    """Read injected violation metadata."""
    if not metadata_path.exists():
        return {"violations": []}
    
    try:
        return json.loads(metadata_path.read_text())
    except Exception:
        return {"violations": []}

def analyze_variation(variation_dir: Path) -> Optional[TestResult]:
    """Analyze a single test variation."""
    
    # Read injected violations
    metadata = read_violation_metadata(variation_dir / "violation_metadata.json")
    violations_injected = len(metadata.get("violations", []))
    
    # Read cli-review output
    report_path = variation_dir / "cli-review" / "cli-review.md"
    if not report_path.exists():
        return None
    
    report_content = report_path.read_text()
    
    # Extract detected issues from report
    detected_issues = extract_findings_from_report(report_path)
    total_detected = sum(issue["count"] for issue in detected_issues)
    false_negatives = max(0, violations_injected - total_detected)
    
    # Count severity breakdown of injected violations
    severity_breakdown = {
        "HIGH": sum(1 for v in metadata.get("violations", []) if v.get("severity") == "HIGH"),
        "MEDIUM": sum(1 for v in metadata.get("violations", []) if v.get("severity") == "MEDIUM"),
        "LOW": sum(1 for v in metadata.get("violations", []) if v.get("severity") == "LOW")
    }
    
    return TestResult(
        variation_id=variation_dir.name,
        violations_injected=violations_injected,
        findings_detected=total_detected,
        false_negatives=false_negatives,
        severity_breakdown=severity_breakdown,
        detected_issues=detected_issues
    )

def main():
    test_dir = Path("/project/tests")
    results = []
    
    # Process all variations
    for i in range(1, 21):
        var_dir = test_dir / f"todo-{i:02d}"
        if var_dir.exists():
            result = analyze_variation(var_dir)
            if result:
                results.append(result)
                print(f"✓ Analyzed {result.variation_id}: {result.violations_injected} injected, {result.findings_detected} detected")
            else:
                print(f"✗ No cli-review output for {var_dir.name}")
    
    if not results:
        print("No variations analyzed!")
        return
    
    # Calculate summary statistics
    total_injected = sum(r.violations_injected for r in results)
    total_detected = sum(r.findings_detected for r in results)
    total_false_negatives = sum(r.false_negatives for r in results)
    
    high_injected = sum(r.severity_breakdown["HIGH"] for r in results)
    high_detected = sum(r.detected_issues[0]["count"] if r.detected_issues and r.detected_issues[0]["severity"] == "High" else 0 for r in results)
    
    summary = {
        "harness_version": "2.0",
        "total_variations": len(results),
        "total_violations_injected": total_injected,
        "total_violations_detected": total_detected,
        "total_false_negatives": total_false_negatives,
        "detection_coverage_percent": (total_detected / total_injected * 100) if total_injected > 0 else 0,
        "severity_summary": {
            "HIGH": {
                "injected": high_injected,
                "detected": high_detected,
                "coverage_percent": (high_detected / high_injected * 100) if high_injected > 0 else 0
            }
        },
        "results": [asdict(r) for r in results]
    }
    
    # Write analysis
    output_file = test_dir / "ANALYSIS.json"
    with open(output_file, "w") as f:
        json.dump(summary, f, indent=2)
    
    print(f"\n✓ Analysis saved to {output_file}")
    print(f"\nSummary:")
    print(f"  Total variations: {summary['total_variations']}")
    print(f"  Total injected violations: {summary['total_violations_injected']}")
    print(f"  Total detected violations: {summary['total_violations_detected']}")
    print(f"  False negatives: {summary['total_false_negatives']}")
    print(f"  Detection coverage: {summary['detection_coverage_percent']:.1f}%")

if __name__ == "__main__":
    main()
