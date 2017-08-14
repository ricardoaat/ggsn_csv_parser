# GGSN parsed csv's reporter
---
Goes through the files created by the GGSN's CDR parser in order to generate a daily report. After finishing processing it compress and store the files, and erases the parser's output files.


### Process CDR files
---
Loads the CSV files created by the parser by day for the amount of days setted with the flag -days (20 Days by default if the flag is not used). Adds the file’s fields download, upload, sumbytes in total, by rating group and mvno (using a range of IMSIs configured on the configuration file)

### Results files creation 
---
Creates two files containing the information gathered in the processing of the CDR files:

#### ggsn_parsed_report_[date YYYYMMDD].csv
Example:
```
CDRday|Download|Upload|Sumbytes #Header
20170730|64639928325|7330136159|71970064484 #According to header
RG:0 0|RG:22 669778578|RG:6 166159|RG:3 83074|RG:60 22802533|RG:70 4021885|RG:34 4652887959|RG:10 12330318207|RG:35 2203647590 #Rating group/data relation
leone 2455986|letwo 47822742|lethree 71919785756| #MVNO/data relation
```

#### imsis_without_mvno_[date YYYYMMDD].csv
Contains a list of IMSIs that were not in the imsi ranges per MVNO provided in the configuration file

### Backup process
---
Compress, copy and remove the processed files. In the configuration file _config.toml_ can be configure the remote destination (SFTP) in the SFTPDest under the Path section key.

To deactivate this process the following flag could be set on false in the configuration file.
``` Toml
[Process]
Backup = false 
```
Set whether it should or should not run the backup process

## FLAGS
---
	-v Shows build and version information
	-days Number of days in the past to check for CDR files, default 20


## BUILD
---
In order to compile the GO binary change the DEPLOY_PATH on the Makefile and run

	make build


## Author
---
Ricardo Arcila