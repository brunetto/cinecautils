#!/usr/bin/env python

import os, time

# Credits: https://www.youtube.com/watch?v=z13qnzUQwuI

t_init = time.time()

# Path to touch, sorry, you need to write them by hand for now
cinecaHome = "/eurora/home/userexternal/bziosi00"
cinecaScratch = "/gpfs/scratch/userexternal/bziosi00"

# Loop on home and scratch
for rootFolder in [cinecaHome, cinecaScratch]:
	# List to keep track of the absolute path, if not, 
	# python can't find the files
	paths = []
	# Loop on the folders and files
	for folder, subfolders, files in os.walk(rootFolder):
		# Keep track of the path
		paths.append(folder)
		# Touch the folder
		os.utime(folder, None)
		# Touch the files
		for f in files:
			os.utime(os.path.join(*(paths + [f])), None)
			
# If everything worked, write a log line to know,
# if not, maybe the script was killed because I
# was too lazy to submit it to a queue 
# and he was too slow to finish in time
f = open("touched.txt", "a")
f.write('Done on ' + str(time.strftime("%Y/%m/%d, %H:%M:%S")) + ' in ' + str(time.time() - t_init) + ' secs \n') 
f.flush()
f.close()


