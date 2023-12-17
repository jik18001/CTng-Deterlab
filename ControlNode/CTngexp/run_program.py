import subprocess
import time
import sys

def run_program(executable_path, role, index):
    cmd = [executable_path, role, index]
    p = subprocess.Popen(cmd)
    time.sleep(300)  # Sleep for 300 seconds after the process has started
    print(f"Terminating {role} {index}...")
    p.terminate()

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python run_program.py <role> <index>")
        sys.exit(1)
    run_program("./test1_executable", sys.argv[1], sys.argv[2])
