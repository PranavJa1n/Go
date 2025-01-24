from flask import Flask
app = Flask(__name__)

@app.route("/")
def inn():
    return """
<!DOCTYPE html>
<html>
<head>
    <title> hello </title>
</head>
<body>
    <h1> HELLO </h1><br>
    idsbvifbvifeonvoefnvoe<br>
    vefiuvieoubvolwc<br>
    vbuowdncds
</body>
</html>
"""

if __name__ == "__main__":
    app.run()