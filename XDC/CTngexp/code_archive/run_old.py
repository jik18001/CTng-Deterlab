import subprocess
import time

def run_programs():
    executable_path = "./test1_executable"
    
    # Define all the commands to run
    commands = [
        [executable_path, "Logger", "1"],
        [executable_path, "Logger", "2"],
        [executable_path, "Logger", "3"],
        [executable_path, "Logger", "4"],
        [executable_path, "CA", "1"],
        [executable_path, "CA", "2"],
        [executable_path, "CA", "3"],
        [executable_path, "CA", "4"],
        [executable_path, "Monitor", "1"],
        [executable_path, "Monitor", "2"],
        [executable_path, "Monitor", "3"],
        [executable_path, "Monitor", "4"],
        [executable_path, "Gossiper", "1"],
        [executable_path, "Gossiper", "2"],
        [executable_path, "Gossiper", "3"],
        [executable_path, "Gossiper", "4"],
    ]

    # Launch all the commands with a slight delay
    processes = []
    for cmd in commands:
        p = subprocess.Popen(cmd)
        processes.append(p)
        time.sleep(0.1)  # Sleep for 0.1 seconds before starting the next process

    # Sleep for 300 seconds after all processes have started
    time.sleep(300)

    # Run Go tests (if you have tests written separately, you can run those too. I'm omitting it for now.)
    print("Terminating all processes...")
    for p in processes:
        p.terminate()

    print("All processes terminated.")


if __name__ == "__main__":
    run_programs()
