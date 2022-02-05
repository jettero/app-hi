
# Wassis?

I wrote an app hundred years ago
[App::HI](https://github.com/jettero/term--ansicolorx) in Perl that helped me
colorize stream text and log text for the purposes of tracking things (mac
addresses, keywords, ip addresses, etc) over time or through the text.

It's particularly useful for things like:

    tcpdump -vlnei blah0 "(complicated filter thingy)" \
        | hi . coal addr:1 lime addr:2 '\.80\b' red '\.443\b' green

now as the packets roll in, the two addresses and a couple of port numbers are
going to be instantly visible in the streaming wall of packet info.

# Whythisthen?

As the years increment since I stopped writing Perl — not out of spite or malace
mind you … I still love Perl, I just don't use it at work like ever — it's
becoming more and more difficult, time consuming and/or annoying to install a
working perl, CPAN::Minus, and finally App::HI everywhere I go.

At my current job, I think it'd be essentially impossible to get it installed on
even a tiny faction of the hosts where I'd like to use it …

… but man do they love Golang at this place …

And Golang has the distinct advantage that you can build a binary once and just
copy it over to the million hosts you want to use it on and you're done. Works.
Nothing else to do.
