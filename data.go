package main

type Pericopes struct {
	Date    string
	Psalm   []string
	First   []string
	Epistle []string
	Gospel  []string
}

var (
	PericopeData []Pericopes = []Pericopes{
		{"2014-11-30", []string{"Psalm 80:1-7"}, []string{"Isaiah 64:1-9"}, []string{"1 Corinthians 1:3-9"}, []string{"Mark 11:1-10", "Mark 13:24-37"}},
		{"2014-12-07", []string{"Psalm 85"}, []string{"Isaiah 40:1-11"}, []string{"2 Peter 3:8-14"}, []string{"Mark 1:1-8"}},
		{"2014-12-14", []string{"Psalm 126"}, []string{"Isaiah 61:1-4", "Isaiah 61:8-11"}, []string{"1 Thessalonians 5:16-24"}, []string{"John 1:6-8", "John 1:19-28"}},
		{"2014-12-21", []string{"Psalm 89:1-5", "Psalm 89:19-29"}, []string{"2 Samuel 7:1-11", "2 Samuel 7:16"}, []string{"Romans 16:25-27"}, []string{"Luke 1:26-38"}},
		{"2014-12-24", []string{"Psalm 110:1-4"}, []string{"Isaiah 7:10-14"}, []string{"1 John 4:7-16"}, []string{"Matthew 1:18-25"}},
		{"2014-12-24", []string{"Psalm 96"}, []string{"Isaiah 9:2-7"}, []string{"Titus 2:11-14"}, []string{"Luke 2:1-20"}},
		{"2014-12-25", []string{"Psalm 98"}, []string{"Isaiah 62:10-12"}, []string{"Titus 3:4-7"}, []string{"Luke 2:1-20"}},
		{"2014-12-25", []string{"Psalm 2"}, []string{"Isaiah 52:7-10"}, []string{"Hebrews 1:1-12"}, []string{"John 1:1-18"}},
		{"2014-12-28", []string{"Psalm 111"}, []string{"Isaiah 61:10—62:3"}, []string{"Galatians 4:4-7"}, []string{"Luke 2:22-40"}},
		{"2014-12-31", []string{"Psalm 90:1-12"}, []string{"Isaiah 30:8-17"}, []string{"Romans 8:31-39"}, []string{"Luke 12:35-40"}},
		{"2015-01-01", []string{"Psalm 8"}, []string{"Numbers 6:22-27"}, []string{"Galatians 3:23-29"}, []string{"Luke 2:21"}},
		{"2015-01-04", []string{"Psalm 119:97-104"}, []string{"1 Kings 3:4-15"}, []string{"Ephesians 1:3-14"}, []string{"Luke 2:40-52"}},
		{"2015-01-06", []string{"Psalm 72:1-15"}, []string{"Isaiah 60:1-6"}, []string{"Ephesians 3:1-12"}, []string{"Matthew 2:1-12"}},
		{"2015-01-11", []string{"Psalm 29"}, []string{"Genesis 1:1-5"}, []string{"Romans 6:1-11"}, []string{"Mark 1:4-11"}},
		{"2015-01-18", []string{"Psalm 139:1-10"}, []string{"1 Samuel 3:1-20"}, []string{"1 Corinthians 6:12-20"}, []string{"John 1:43-51"}},
		{"2015-01-25", []string{"Psalm 62"}, []string{"Jonah 3:1-5", "Jonah 3:10"}, []string{"1 Corinthians 7:29-35"}, []string{"Mark 1:14-20"}},
		{"2015-02-01", []string{"Psalm 111"}, []string{"Deuteronomy 18:15-20"}, []string{"1 Corinthians 8:1-13"}, []string{"Mark 1:21-28"}},
		{"2015-02-08", []string{"Psalm 147:1-11"}, []string{"Isaiah 40:21-31"}, []string{"1 Corinthians 9:16-27"}, []string{"Mark 1:29-39"}},
		{"2015-02-15", []string{"Psalm 50:1-6"}, []string{"2 Kings 2:1-12", "Exodus 34:29-35"}, []string{"2 Corinthians 3:12-18", "2 Corinthians 4:1-6"}, []string{"Mark 9:2-9"}},
		{"2015-02-18", []string{"Psalm 51:1-19"}, []string{"Joel 2:12-19"}, []string{"2 Corinthians 5:20—6:10"}, []string{"Matthew 6:1-6", "Matthew 6:16-21"}},
		{"2015-02-22", []string{"Psalm 25:1-10"}, []string{"Genesis 22:1-18"}, []string{"James 1:12-18"}, []string{"Mark 1:9-15"}},
		{"2015-03-01", []string{"Psalm 22:23-31"}, []string{"Genesis 17:1-7", "Genesis 17:15-16"}, []string{"Romans 5:1-11"}, []string{"Mark 8:27-38"}},
		{"2015-03-08", []string{"Psalm 19"}, []string{"Exodus 20:1-17"}, []string{"1 Corinthians 1:18-31"}, []string{"John 2:13-25"}},
		{"2015-03-15", []string{"Psalm 107:1-9"}, []string{"Numbers 21:4-9"}, []string{"Ephesians 2:1-10"}, []string{"John 3:14-21"}},
		{"2015-03-22", []string{"Psalm 119:9-16"}, []string{"Jeremiah 31:31-34"}, []string{"Hebrews 5:1-10"}, []string{"Mark 10:32-45"}},
		{"2015-03-29", []string{"Psalm 118:19-29", "Psalm 31:9-16"}, []string{"Zechariah 9:9-12"}, []string{"Philippians 2:5-11"}, []string{"Mark 14:1—15:47", "Mark 15:1-47", "John 12:20-43"}},
		{"2015-04-02", []string{"Psalm 116:12-19"}, []string{"Exodus 24:3-11"}, []string{"1 Corinthians 10:16-17"}, []string{"Mark 14:12-26"}},
		{"2015-04-03", []string{"Psalm 22", "Psalm 31"}, []string{"Isaiah 52:13—53:12"}, []string{"Hebrews 4:14-16", "Hebrews 5:7-9"}, []string{"John 18:1—19:42", "John 19:17-30"}},
		{"2015-04-05", []string{"Psalm 118:15-29"}, []string{"Exodus 15:1-11"}, []string{"1 Corinthians 5:6-8"}, []string{"John 20:1-18"}},
		{"2015-04-05", []string{"Psalm 16"}, []string{"Isaiah 25:6-9"}, []string{"1 Corinthians 15:1-11"}, []string{"Mark 16:1-8"}},
		{"2015-04-12", []string{"Psalm 148"}, []string{"Acts 4:32-35"}, []string{"1 John 1:1—2:2"}, []string{"John 20:19-31"}},
		{"2015-04-19", []string{"Psalm 4"}, []string{"Acts 3:11-21"}, []string{"1 John 3:1-7"}, []string{"Luke 24:36-49"}},
		{"2015-04-26", []string{"Psalm 23"}, []string{"Acts 4:1-12"}, []string{"1 John 3:16-24"}, []string{"John 10:11-18"}},
		{"2015-05-03", []string{"Psalm 150"}, []string{"Acts 8:26-40"}, []string{"1 John 4:1-21"}, []string{"John 15:1-8"}},
		{"2015-05-10", []string{"Psalm 98"}, []string{"Acts 10:34-48"}, []string{"1 John 5:1-8"}, []string{"John 15:9-17"}},
		{"2015-05-14", []string{"Psalm 47"}, []string{"Acts 1:1-11"}, []string{"Ephesians 1:15-23"}, []string{"Luke 24:44-53"}},
		{"2015-05-17", []string{"Psalm 1"}, []string{"Acts 1:12-26"}, []string{"1 John 5:9-15"}, []string{"John 17:11-19"}},
		{"2015-05-24", []string{"Psalm 139:1-16"}, []string{"Ezekiel 37:1-14"}, []string{"Acts 2:1-21"}, []string{"John 15:26-27", "John 16:4-15"}},
		{"2015-05-31", []string{"Psalm 29"}, []string{"Isaiah 6:1-8"}, []string{"Acts 2:14", "Acts 2:22-36"}, []string{"John 3:1-17"}},
		{"2015-06-07", []string{"Psalm 130"}, []string{"Genesis 3:8-15"}, []string{"2 Corinthians 4:13—5:1"}, []string{"Mark 3:20-35"}},
		{"2015-06-14", []string{"Psalm 1"}, []string{"Ezekiel 17:22-24"}, []string{"2 Corinthians 5:1-17"}, []string{"Mark 4:26-34"}},
		{"2015-06-21", []string{"Psalm 124"}, []string{"Job 38:1-11"}, []string{"2 Corinthians 6:1-13"}, []string{"Mark 4:35-41"}},
		{"2015-06-28", []string{"Psalm 30"}, []string{"Lamentations 3:22-33"}, []string{"2 Corinthians 8:1-9", "2 Corinthians 8:13-15"}, []string{"Mark 5:21-43"}},
		{"2015-07-05", []string{"Psalm 123"}, []string{"Ezekiel 2:1-5"}, []string{"2 Corinthians 12:1-10"}, []string{"Mark 6:1-13"}},
		{"2015-07-12", []string{"Psalm 85:1-13"}, []string{"Amos 7:7-15"}, []string{"Ephesians 1:3-14"}, []string{"Mark 6:14-29"}},
		{"2015-07-19", []string{"Psalm 23"}, []string{"Jeremiah 23:1-6"}, []string{"Ephesians 2:11-22"}, []string{"Mark 6:30-44"}},
		{"2015-07-26", []string{"Psalm 136:1-9"}, []string{"Genesis 9:8-17"}, []string{"Ephesians 3:14-21"}, []string{"Mark 6:45-56"}},
		{"2015-08-02", []string{"Psalm 145:10-21"}, []string{"Exodus 16:2-15"}, []string{"Ephesians 4:1-16"}, []string{"John 6:22-35"}},
		{"2015-08-09", []string{"Psalm 34:1-8"}, []string{"1 Kings 19:1-8"}, []string{"Ephesians 4:17—5:2"}, []string{"John 6:35-51"}},
		{"2015-08-16", []string{"Psalm 34:12-22"}, []string{"Proverbs 9:1-10", "Joshua 24:1-2", "Joshua 24:14-18"}, []string{"Ephesians 5:6-21"}, []string{"John 6:51-69"}},
		{"2015-08-23", []string{"Psalm 14"}, []string{"Isaiah 29:11-19"}, []string{"Ephesians 5:22-33"}, []string{"Mark 7:1-13"}},
		{"2015-08-30", []string{"Psalm 119:129-136"}, []string{"Deuteronomy 4:1-2", "Deuteronomy 4:6-9"}, []string{"Ephesians 6:10-20"}, []string{"Mark 7:14-23"}},
		{"2015-09-06", []string{"Psalm 146"}, []string{"Isaiah 35:4-7"}, []string{"James 2:1-10", "James 2:14-18"}, []string{"Mark 7:24-37"}},
		{"2015-09-13", []string{"Psalm 116:1-9"}, []string{"Isaiah 50:4-10"}, []string{"James 3:1-12"}, []string{"Mark 9:14-29"}},
		{"2015-09-20", []string{"Psalm 54"}, []string{"Jeremiah 11:18-20"}, []string{"James 3:13—4:10"}, []string{"Mark 9:30-37"}},
		{"2015-09-27", []string{"Psalm 104:27-35"}, []string{"Numbers 11:4-6", "Numbers 11:10-16", "Numbers 11:24-29"}, []string{"James 5:1-20"}, []string{"Mark 9:38-50"}},
		{"2015-10-04", []string{"Psalm 128"}, []string{"Genesis 2:18-25"}, []string{"Hebrews 2:1-18"}, []string{"Mark 10:2-16"}},
		{"2015-10-11", []string{"Psalm 90:12-17"}, []string{"Amos 5:6-7", "Amos 5:10-15"}, []string{"Hebrews 3:12-19"}, []string{"Mark 10:17-22"}},
		{"2015-10-18", []string{"Psalm 119:9-16"}, []string{"Ecclesiastes 5:10-20"}, []string{"Hebrews 4:1-16"}, []string{"Mark 10:23-31"}},
		{"2015-10-25", []string{"Psalm 46"}, []string{"Revelation 14:6-7"}, []string{"Romans 3:19-28"}, []string{"John 8:31-36", "Matthew 11:12-19"}},
		{"2015-11-01", []string{"Psalm 149"}, []string{"Revelation 7:2-17"}, []string{"1 John 3:1-3"}, []string{"Matthew 5:1-12"}},
		{"2015-11-08", []string{"Psalm 146"}, []string{"1 Kings 17:8-16"}, []string{"Hebrews 9:24-28"}, []string{"Mark 12:38-44"}},
		{"2015-11-15", []string{"Psalm 16"}, []string{"Daniel 12:1-3"}, []string{"Hebrews 10:11-25"}, []string{"Mark 13:1-13"}},
		{"2015-11-22", []string{"Psalm 93"}, []string{"Isaiah 51:4-6", "Daniel 7:9-10", "Daniel 7:13-14"}, []string{"Jude 20-25", "Revelation 1:4-8"}, []string{"Mark 13:24-37", "John 18:33-37"}},
		{"2015-11-26", []string{"Psalm 67"}, []string{"Deuteronomy 8:1-10"}, []string{"Philippians 4:6-20", "1 Timothy 2:1-4"}, []string{"Luke 17:11-19"}},
		{"2015-11-29", []string{"Psalm 25:1-10"}, []string{"Jeremiah 33:14-16"}, []string{"1 Thessalonians 3:9-13"}, []string{"Luke 19:28-40", "Luke 21:25-36"}},
	}
)

func str(slice []string) string {
	for _, s := range slice {
		buffer.WriteString(s)
	}
	concatenated := buffer.String()
	buffer.Reset()
	return concatenated
}

func slc(args ...string) []string {
	return args
}

func combine(args ...string) string {
	return str(args)
}
