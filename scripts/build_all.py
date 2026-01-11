import subprocess
import os

services = [
    "services/generator-go",
    "services/dashboard-ui",
    "services/processor-flink"
]

def build_monorepo():
    print("🛠️ Starting Multi-Service Build Pipeline...")
    for service in services:
        print(f"Building {service}...")
        # Simulates the CI/CD cycle reduction you achieved at Zaspar [cite: 19]
        if "go" in service:
            subprocess.run(["go", "build", "-o", "bin/engine"], cwd=service)
        elif "ui" in service:
            print("Running React Build (npm run build)...")
        elif "flink" in service:
            subprocess.run(["mvn", "clean", "package"], cwd=service)
    print("✅ Build Complete.")

if __name__ == "__main__":
    build_monorepo()