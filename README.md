
# Wassis?

I wrote an app a hundred years ago
[App::HI](https://github.com/jettero/term--ansicolorx)… in Perl… that helped me
colorize stream text and log text for the purposes of tracking things over time
through the text (mac addresses, keywords, ip addresses, and the like).

It's particularly useful for things like:

    tcpdump -lnei vhb.42 -c 3 tcp port 22 \
        hi . coal \
        fe:ed:be:ef:00:00 purple \
        fe:ed:be:ef:00:01 sky \
        '\d+\.\d+\.\d+\.\d+\.\d+' green \
        '\d+\.\d+\.\d+\.\d+' yellow

now as the packets roll in, the two addresses and a couple of port numbers are
going to be instantly visible in the streaming wall of packet info.

![example output](cruft/example.png?raw=true "example output")

# Whythisthen?

As the years increment since I stopped writing Perl — not out of spite or malace
mind you … I still love Perl, I just don't use it at work like ever. It's
becoming more and more difficult, time consuming and/or annoying to install a
working Perl, CPAN::Minus, and App::HI everywhere I go.

At my current job, I think it'd be essentially impossible to get it installed on
even a tiny faction of the hosts where I'd like to use it …

… but man do they love Golang at this place …

And Golang has the distinct advantage that you can build a binary once and just
copy it over to the million hosts you want to use it on and you're done. Works.
Nothing else to do.
