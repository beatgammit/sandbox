Intro
=====

Mobile Safari bug
-----------------

Mobile Safari has a bug (or maybe it is jQuery that has the bug) that reduces functionality of long polling when updating a canvas element.

Included are three example pages:

* `no-timeout.html`- first code example, preferred method (does not work)
* `timeout-in.html`- second example, timeout inside the success callback
* `timeout-out.html`- third example, timeout outside the success callback

Also included is a node.js server, called bugserver, but this isn't anything special.

Long polling call can be made like this:

    function longPoll(url) {
        $.ajax({
            url: url,
            success: function (data) {
                // do stuff with data and draw to canvas
    
                longPoll(url);
            }
        });
    }

This works on most browsers, but not on mobile safari when run as a home screen web app.  It works fine on mobile safari in the regular browser though.

A workaround can be made like this:

    function longPoll(url) {
        $.ajax({
            url: url,
            success: function (data) {
                // do stuff with data and draw to canvas
    
                setTimeout(function () {
                    longPoll(url);
                }, 50);
            }
        });
    }

Or like this:

    function longPoll(url) {
        $.ajax({
            url: url,
            success: function (data) {
                // do stuff with data and draw to canvas
            }
        });
        
        setTimeout(function () {
            longPoll(url);
        }, 50);
    }

The weird thing, though, is that the ajax requests continue to be made and received successfully, but the canvas is not updated.  From what I can tell, this only applies to the canvas element.

This is most likely not a problem with WebKit because:

* It works fine in the regular mobile Safari
* It works fine on regular Safari
* It works fine on Chrome
* It works fine the Android browser

It seems that Apple uses different code in its 'home-screen' browser than in its regular Safari browser, which makes absolutely no sense.
