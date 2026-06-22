import express from 'express';

const app = express();

app.get('/', (req, res) => {
    res.send('Hello World!');
});

app.listen(8006, '0.0.0.0', () => {
    console.log('listening on 8006')
});