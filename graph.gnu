set title "Frequency of Money Winnings (1 Billion Trials)"
show title
set xlabel "Amount Won ($)"
set ylabel "Frequency of Winning"
show xlabel
show ylabel
plot "data/graph.dat" using 1:2 with lines