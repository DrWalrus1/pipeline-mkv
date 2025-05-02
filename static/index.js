let handleInfoEvent = event => {
    console.log(JSON.parse(event.data))
    // Handle the message data here.  For example, display it in a div:
    const messageDiv = document.getElementById('messages');
    if (messageDiv) {
      messageDiv.innerHTML += `<p>${event.data}</p>`;
    }
};

function connectWS(action, messageHandler) {
  const socket = new WebSocket(`ws://${window.location.hostname}:8080/api/${action}?source=disc:0`); // Replace with your actual WebSocket URL

  socket.addEventListener('open', (event) => {
    console.log('WebSocket connection opened:', event);
  });

  socket.addEventListener('message', messageHandler);
  socket.addEventListener('close', (event) => {
    console.log('WebSocket connection closed:', event);
  });

  socket.addEventListener('error', (error) => {
    console.error('WebSocket error:', error);
  });

  // Example: Sending a message to the server (e.g., to cancel)
  const cancelButton = document.getElementById('cancelButton'); // Assuming you have a button with this ID
  if (cancelButton) {
    cancelButton.addEventListener('click', () => {
      socket.send('cancel');
      console.log('Cancel message sent');
    });
  }
}


