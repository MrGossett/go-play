var ws = new WebSocket("ws://localhost:7000/conn");
ws.onmessage = function(e) { console.log(e.data); };

ws.send(JSON.stringify({method: "Arith.Multiply", params: [{"A": 4, "B": 5}], id: 1}));
ws.send(JSON.stringify({method: "Arith.Average", params: [[1,2,3,4,5,6,7,8,9]], id: 2}));
