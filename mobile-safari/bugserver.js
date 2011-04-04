var connect = require('connect'),
    port = 8888;

connect(
    connect.static('./')
).listen(port, function () {
    console.log("Listening on port " + port);
});
