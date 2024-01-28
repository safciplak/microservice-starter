const apimock       = require('@ng-apimock/core');
const devInterface  = require('@ng-apimock/dev-interface');
const express       = require('express');
const morgan        = require('morgan');

const app = express();

// Enable logging to the terminal
app.use(morgan('dev'));

// Set the port to 3000 or the supplied PORT in the environment
app.set('port', (process.env.PORT || 3000));

// Process the application mocks
apimock.processor.process({src: './v1'});

// Add mocking as middleware ot the express server
app.use(apimock.middleware);

// Bind a route so you can modify the running mock server
app.use('/mocking', express.static(devInterface));

app.listen(app.get('port'), function() {
    console.log('app running on port', app.get('port'));
});