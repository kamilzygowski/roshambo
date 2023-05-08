import * as io  from 'socket.io-client';

// "undefined" means the URL will be computed from the `window.location` object
const URL = "http://localhost:8000/socket.io";

export const socket = io(URL, {
    //transports: ["websocket"],
    transports: ["polling"],
    reconnection: true
});