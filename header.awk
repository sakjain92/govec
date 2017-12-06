BEGIN{
	count = 0;
	ORS = ""
}
{
	if(NF > 0) {
		if($1 == "export") {
			for(i = 1; i <= NF; i++) {
				len = split($i, tokens, "(")
				for(j = 1; j <= len; j++) {
					if(tokens[j] != "export" && tokens[j] != "uniform" && tokens[j] != "{") {
						print tokens[j]
						print " "
					}
					if(j != 1) {
						print "("
					}
				}
			} 
			print ";\n"
		}
	}
	count += 1;
}
END {
	print count
}