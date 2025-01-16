from flask import Flask, jsonify, request, send_from_directory
import os

app = Flask(__name__)

# Mock data for events
events = {
    "1": {"name": "Rock Concert", "tickets_available": 100},
    "2": {"name": "Tech Conference", "tickets_available": 50},
    "3": {"name": "Art Exhibition", "tickets_available": 20}
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
    app.run(host="http://rock-concert-github-test-1.apps.67853a4ffeba7dc9c41f3374.ocp.techzone.ibm.com", port=8080)
