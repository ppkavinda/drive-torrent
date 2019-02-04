import {EventEmitter} from 'events';

class Socket {
    constructor(wsurl, ee = new EventEmitter()) {
        let ws = new WebSocket(wsurl);
        this.ee = ee;
        this.ws = ws;
        ws.onmessage = this.message.bind(this);
        ws.onopen = this.open.bind(this);
        ws.onclose = this.close.bind(this);
    }
    on(name, fn) {
        this.ee.on(name, fn);
    }
    off(name, fn) {
        this.ee.removeListener(name, fn);
    }
    emit(name, data) {
        const message = JSON.stringify({name, data});
        this.ws.send(message);
    }
    message(e) {
        try {
            const msgData = JSON.parse(e.data);
            this.ee.emit(msgData.Event, msgData.Data);
        }
        catch(err) {
            let error = {
                message: err
            }
            console.log(err)
            this.ee.emit(error.message)
        }
    }
    open() {
        this.ee.emit('connected');
    }
    close() {
        this.ee.emit('disconnected');
    }   
}

export default Socket