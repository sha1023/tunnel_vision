# tunnel_vision

## Idea
There is a problem a lot of writers have of editing old sentences when they should be writing new sentences. A popular trick is to have a small window that prevents them from seeing what they have already written.

This is meant to provide such an interface. In theory it's meant to be portable across operating systems because it's run on a browser. My front end chops are way too small too guarantee anything. (Also I've used the oldest most rudimenttary js and html that I know of, so it might already be worse than obsolete.)

## Installation/Use

Just put the tunnel_vision.html on your desktop and open it with your favorite browser. Probably double clicking it should work.

Use ought to be self explanatory. You can resize the window with the mouse and set the amount of text you can see via the memory length input. Clicking the "What did I write?" button displays what you wrote. You have to copy and paste it into a file to save it if you click back, you lose your work. (Sorry guys, might improve this later.)

## WARNING

Ideally it would save to a file, but embedded javascript is generally prevented from doing that. Instead I write the saved words to a string, which needs to be copy pasted by the user, also if you close the window or click back you lose your work. (Extremely kluge I know.) I'm not sure if js string lengths are standardized but on my machine the memory is big enough to fit at least 70000 words. (Try pasting http://www.gutenberg.org/cache/epub/30165/pg30165.txt into the window and see if it is properly saved.)
