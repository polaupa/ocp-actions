from flask import Flask, jsonify, request, send_from_directory
from flask_cors import CORS
import os

app = Flask(__name__)
CORS(app)

events = {
    "1": {"name": "Rock Concert", "tickets_available": 20000},
    "2": {"name": "Tech Conference", "tickets_available": 5000},
    "3": {"name": "Art Exhibition", "tickets_available": 100}
}

# Serve the frontend file
@app.route("/")
def serve_frontend():
    return send_from_directory(os.getcwd(), "index.html")

# API to get events
@app.route("/events", methods=["GET"])
def get_events():
    return jsonify(events)

# API to purchase tickets
@app.route("/events/<event_id>/purchase", methods=["POST"])
def purchase_tickets(event_id):
    data = request.json
    tickets_requested = data.get("tickets", 0)

    if event_id not in events:
        return jsonify({"error": "Event not found"}), 404

    if events[event_id]["tickets_available"] < tickets_requested:
        return jsonify({"error": "Not enough tickets available"}), 400

    events[event_id]["tickets_available"] -= tickets_requested
    return jsonify({"message": f"Successfully purchased {tickets_requested} tickets for {events[event_id]['name']}!"})

# Run the app
if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)

