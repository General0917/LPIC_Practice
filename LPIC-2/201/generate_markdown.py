
import os

def create_markdown_file(base_filename="numbered_list.md", start=1, end=60):
    """
    Generates a Markdown file with a numbered list, avoiding overwriting existing files.

    Args:
        base_filename (str): The base name of the Markdown file to create.
        start (int): The starting number of the list.
        end (int): The ending number of the list.
    """
    filename = base_filename
    counter = 1
    name, ext = os.path.splitext(filename)
    while os.path.exists(filename):
        filename = f"{name}({counter}){ext}"
        counter += 1

    try:
        with open(filename, 'w', encoding='utf-8') as f:
            for i in range(start, end + 1):
                f.write(f"{i}. \n")
        print(f"Successfully created '{filename}' with numbers from {start} to {end}.")
    except IOError as e:
        print(f"Error writing to file {filename}: {e}")

if __name__ == "__main__":
    output_filename = "課題リスト.md"
    create_markdown_file(output_filename, 1, 60)
    # Print current working directory to show where the file is created
    print(f"File created in: {os.getcwd()}")
