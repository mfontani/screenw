#!/usr/bin/env perl
use 5.020_001;
use warnings;
use Term::ReadKey qw<GetTerminalSize>;
my ($width) = GetTerminalSize();
my @lines = (
    '',
    '',
    '',
);
for my $n (1..$width) {
    $lines[0] .= $n % 100 == 0 ? int($n/100) % 100 : ' ';
    $lines[1] .= $n % 10  == 0 ? int($n/10)  % 10  : ' ';
    $lines[2] .= $n % 10;
}
@lines = grep { $_ =~ /\S/xms } @lines;
@lines = map { sprintf "%-${width}s", $_ } @lines;
$width = "=$width";
substr $lines[0], - length $width, length $width, $width;
say for @lines;
