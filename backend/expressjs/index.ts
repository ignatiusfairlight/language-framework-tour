import express from 'express';
import routes from './src/routes';

const app = express();

app.use(express.json());
app.use(routes);

app.listen(8006, '0.0.0.0', () => {
    console.log('listening on 8006')
});