import subprocess
import sys

def deploy_stack(stack_name):
    """Automates AWS CloudFormation deployment [cite: 12, 29]"""
    print(f"--- Deploying {stack_name} Infrastructure ---")
    try:
        cmd = [
            "aws", "cloudformation", "deploy",
            "--template-file", "infrastructure/cloudformation/kinesis-stream.yaml",
            "--stack-name", stack_name,
            "--capabilities", "CAPABILITY_IAM"
        ]
        subprocess.check_call(cmd)
        print("✅ Infrastructure deployed successfully.")
    except Exception as e:
        print(f"❌ Deployment failed: {e}")
        sys.exit(1)

if __name__ == "__main__":
    deploy_stack("finpulse-prod-stack")