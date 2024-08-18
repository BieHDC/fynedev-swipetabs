# fynedev-swipetabs
Experimenting with Swipe Gestures for Fyne  
  
This is a simple test application for quick testing of the thing.  
It also redirects fyne to my repository which has changes to AppTabs  
Right now i am using AppTabs for simplicity, because the basetabs stuff  
is not exported. In the future it would make sense to make an exported  
version of basetabs for this scenario.  
Another issue right now is that there is no smooth scrolling between tabs  
as you would expect in mobile apps. i would prefer to also make a with and  
without animation path for different purposes. however this one would require  
adapting something between the scroll container and tab switcher which also  
would mean we need to keep adjacent tabs "visible" so they are rendered  
before we switch or they are partially visible. should be able to be  
optimised with on drag started.  

Please give feedback and/or contribute. once this is in beta we can  
move to upstreaming.

I recommend running it like `go run -tags debug .` and the repo is set  
up for app building for mobile testing.

## Priority
- debug overlay for insets
- - counldnt get them working taking fysion as draft
- different swipe techiques
- - right now its edge to outside edge
- - but from different sample apps, the swipes
- - can start everywhere, but need to make a
- - specific minimum delta. however for those to
- - look good, we would need a seamless animation.
- - https://stackoverflow.com/questions/15084675/how-to-implement-swipe-gestures-for-mobile-devices/58719294#58719294
