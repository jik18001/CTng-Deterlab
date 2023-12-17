import os
import json
import statistics
from datetime import datetime

# Global variable to track the first call to write_and_print
first_call = True

def write_and_print(file, message):
    """Writes a message to a file and prints it to the console."""
    global first_call
    if first_call:
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        file.write(f"\n--- {timestamp} ---\n")
        first_call = False
    print(message)
    file.write(message + '\n')

def main():
    # Root directory where the Monitor folders are
    root_dir = './'
    output_file_path = 'network_statistics.txt'

    # Initialize empty lists to collect statistics
    all_converge_times = []
    all_converge_times_init = []
    all_total_traffic_received = []
    all_total_traffic_sent = []

    # Open the output file in append mode
    with open(output_file_path, 'a') as output_file:
        # Loop through subfolders (Monitor1, Monitor2, etc.)
        for folder_name in os.listdir(root_dir):
            if "Monitor" in folder_name:
                json_file_path = os.path.join(root_dir, folder_name, 'gossiper_testdata.json')
                if os.path.exists(json_file_path):
                    with open(json_file_path, 'r') as f:
                        data = json.load(f)
                        for key in data:
                            converge_time = float(data[key]['converge_time'])
                            converge_time_init = float(data[key]['converge_time_init'])
                            total_traffic_received = float(data[key]['total_traffic_received'])
                            total_traffic_sent = float(data[key]['total_traffic_sent'])

                            all_converge_times.append(converge_time)
                            all_converge_times_init.append(converge_time_init)
                            all_total_traffic_received.append(total_traffic_received)
                            all_total_traffic_sent.append(total_traffic_sent)
                            break

        if not all_converge_times:
            write_and_print(output_file, "No data found.")
            return

        # Calculate and write/print statistics
        write_and_print(output_file, "Converge Time Statistics:")
        write_and_print(output_file, "Max Converge Time: " + str(max(all_converge_times)))
        write_and_print(output_file, "Average Converge Time: " + str(statistics.mean(all_converge_times)))

        write_and_print(output_file, "Converge Time Init Statistics:")
        write_and_print(output_file, "Max Converge Time Init: " + str(max(all_converge_times_init)))
        write_and_print(output_file, "Average Converge Time Init: " + str(statistics.mean(all_converge_times_init)))

        write_and_print(output_file, "Total Traffic Received Statistics:")
        write_and_print(output_file, "Max Total Traffic Received: " + str(max(all_total_traffic_received)))
        write_and_print(output_file, "Average Total Traffic Received: " + str(statistics.mean(all_total_traffic_received)))

        write_and_print(output_file, "Total Traffic Sent Statistics:")
        write_and_print(output_file, "Max Total Traffic Sent: " + str(max(all_total_traffic_sent)))
        write_and_print(output_file, "Average Total Traffic Sent: " + str(statistics.mean(all_total_traffic_sent)))

if __name__ == "__main__":
    main()
