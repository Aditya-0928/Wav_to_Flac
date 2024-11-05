import base64

# Specify the path to your WAV file
file_path = "C:\\Users\\Aditya\\Downloads\\file_example_WAV_1MG.wav"

# Read the file and encode it in Base64
with open(file_path, "rb") as wav_file:
    base64_data = base64.b64encode(wav_file.read()).decode("utf-8")
    print(base64_data)
