import base64

# Example: your base64 encoded string (replace this with your actual Base64 string)

# Decode the Base64 string
audio_data = base64.b64decode(base64_string)

# Save as a FLAC file
with open("output.flac", "wb") as file:
    file.write(audio_data)

print("FLAC file saved as output.flac")