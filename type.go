package main

import (
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

type Readings struct {
	Psalm   string
	First   string
	Epistle string
	Gospel  string
}

func (r Readings) Slc() []string {
	return []string{r.Psalm, r.First, r.Epistle, r.Gospel}
}

func (r Readings) List() []string {
	s := r.Str()
	return strings.Split(s, ", ")
}

func (r Readings) Str() string {
	slc := r.Slc()
	return ConcatBy(slc, ", ")
}

func (r Readings) Browse(urlBase string) {
	url := r.Url(urlBase)
	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		_ = exec.Command("open", url).Start()
	}
}

func (r Readings) Url(base string) string {
	return Concat(base, r.UrlFmt())
}

func (r Readings) UrlFmt() string {
	return url.QueryEscape(r.Str())
}

type Pericope struct {
	Date     string
	Readings Readings
}

func (p Pericope) Str() string {
	d := p.Date
	r := p.Readings

	s := r.Str()
	return Concat(d, ") ", s)
	//"2014-11-30) Psalm 80:1-7, Isaiah 64:1-9, 1 Corinthians 1:3-9, Mark 11:1-10, Mark 13:24-37"
}

var (
	Lectionary [66]Pericope = [66]Pericope{
		{"2014-11-30",
			Readings{
				"Psalm 80:1-7",
				"Isaiah 64:1-9",
				"1 Corinthians 1:3-9",
				"Mark 11:1-10, Mark 13:24-37"},
		},
		{"2014-12-07",
			Readings{
				"Psalm 85",
				"Isaiah 40:1-11",
				"2 Peter 3:8-14",
				"Mark 1:1-8"},
		},
		{"2014-12-14",
			Readings{
				"Psalm 126",
				"Isaiah 61:1-4, Isaiah 61:8-11",
				"1 Thessalonians 5:16-24",
				"John 1:6-8, John 1:19-28"},
		},
		{"2014-12-21",
			Readings{
				"Psalm 89:1-5, Psalm 89:19-29",
				"2 Samuel 7:1-11, 2 Samuel 7:16",
				"Romans 16:25-27",
				"Luke 1:26-38"},
		},
		{"2014-12-24",
			Readings{
				"Psalm 110:1-4",
				"Isaiah 7:10-14",
				"1 John 4:7-16",
				"Matthew 1:18-25"},
		},
		{"2014-12-24",
			Readings{
				"Psalm 96",
				"Isaiah 9:2-7",
				"Titus 2:11-14",
				"Luke 2:1-20"},
		},
		{"2014-12-25",
			Readings{
				"Psalm 98",
				"Isaiah 62:10-12",
				"Titus 3:4-7",
				"Luke 2:1-20"},
		},
		{"2014-12-25",
			Readings{
				"Psalm 2",
				"Isaiah 52:7-10",
				"Hebrews 1:1-12",
				"John 1:1-18"},
		},
		{"2014-12-28",
			Readings{
				"Psalm 111",
				"Isaiah 61:10—62:3",
				"Galatians 4:4-7",
				"Luke 2:22-40"},
		},
		{"2014-12-31",
			Readings{
				"Psalm 90:1-12",
				"Isaiah 30:8-17",
				"Romans 8:31-39",
				"Luke 12:35-40"},
		},
		{"2015-01-01",
			Readings{
				"Psalm 8",
				"Numbers 6:22-27",
				"Galatians 3:23-29",
				"Luke 2:21"},
		},
		{"2015-01-04",
			Readings{
				"Psalm 119:97-104",
				"1 Kings 3:4-15",
				"Ephesians 1:3-14",
				"Luke 2:40-52"},
		},
		{"2015-01-06",
			Readings{
				"Psalm 72:1-15",
				"Isaiah 60:1-6",
				"Ephesians 3:1-12",
				"Matthew 2:1-12"},
		},
		{"2015-01-11",
			Readings{
				"Psalm 29",
				"Genesis 1:1-5",
				"Romans 6:1-11",
				"Mark 1:4-11"},
		},
		{"2015-01-18",
			Readings{
				"Psalm 139:1-10",
				"1 Samuel 3:1-20",
				"1 Corinthians 6:12-20",
				"John 1:43-51"},
		},
		{"2015-01-25",
			Readings{
				"Psalm 62",
				"Jonah 3:1-5, Jonah 3:10",
				"1 Corinthians 7:29-35",
				"Mark 1:14-20"},
		},
		{"2015-02-01",
			Readings{
				"Psalm 111",
				"Deuteronomy 18:15-20",
				"1 Corinthians 8:1-13",
				"Mark 1:21-28"},
		},
		{"2015-02-08",
			Readings{
				"Psalm 147:1-11",
				"Isaiah 40:21-31",
				"1 Corinthians 9:16-27",
				"Mark 1:29-39"},
		},
		{"2015-02-15",
			Readings{
				"Psalm 50:1-6",
				"2 Kings 2:1-12, Exodus 34:29-35",
				"2 Corinthians 3:12-18, 2 Corinthians 4:1-6",
				"Mark 9:2-9"},
		},
		{"2015-02-18",
			Readings{
				"Psalm 51:1-19",
				"Joel 2:12-19",
				"2 Corinthians 5:20—6:10",
				"Matthew 6:1-6, Matthew 6:16-21"},
		},
		{"2015-02-22",
			Readings{
				"Psalm 25:1-10",
				"Genesis 22:1-18",
				"James 1:12-18",
				"Mark 1:9-15"},
		},
		{"2015-03-01",
			Readings{
				"Psalm 22:23-31",
				"Genesis 17:1-7, Genesis 17:15-16",
				"Romans 5:1-11",
				"Mark 8:27-38"},
		},
		{"2015-03-08",
			Readings{
				"Psalm 19",
				"Exodus 20:1-17",
				"1 Corinthians 1:18-31",
				"John 2:13-25"},
		},
		{"2015-03-15",
			Readings{
				"Psalm 107:1-9",
				"Numbers 21:4-9",
				"Ephesians 2:1-10",
				"John 3:14-21"},
		},
		{"2015-03-22",
			Readings{
				"Psalm 119:9-16",
				"Jeremiah 31:31-34",
				"Hebrews 5:1-10",
				"Mark 10:32-45"},
		},
		{"2015-03-29",
			Readings{
				"Psalm 118:19-29, Psalm 31:9-16",
				"Zechariah 9:9-12",
				"Philippians 2:5-11",
				"Mark 14:1—15:47, Mark 15:1-47, John 12:20-43"},
		},
		{"2015-04-02",
			Readings{
				"Psalm 116:12-19",
				"Exodus 24:3-11",
				"1 Corinthians 10:16-17",
				"Mark 14:12-26"},
		},
		{"2015-04-03",
			Readings{
				"Psalm 22, Psalm 31",
				"Isaiah 52:13—53:12",
				"Hebrews 4:14-16, Hebrews 5:7-9",
				"John 18:1—19:42, John 19:17-30"},
		},
		{"2015-04-05",
			Readings{
				"Psalm 118:15-29",
				"Exodus 15:1-11",
				"1 Corinthians 5:6-8",
				"John 20:1-18"},
		},
		{"2015-04-05",
			Readings{
				"Psalm 16",
				"Isaiah 25:6-9",
				"1 Corinthians 15:1-11",
				"Mark 16:1-8"},
		},
		{"2015-04-12",
			Readings{
				"Psalm 148",
				"Acts 4:32-35",
				"1 John 1:1—2:2",
				"John 20:19-31"},
		},
		{"2015-04-19",
			Readings{
				"Psalm 4",
				"Acts 3:11-21",
				"1 John 3:1-7",
				"Luke 24:36-49"},
		},
		{"2015-04-26",
			Readings{
				"Psalm 23",
				"Acts 4:1-12",
				"1 John 3:16-24",
				"John 10:11-18"},
		},
		{"2015-05-03",
			Readings{
				"Psalm 150",
				"Acts 8:26-40",
				"1 John 4:1-21",
				"John 15:1-8"},
		},
		{"2015-05-10",
			Readings{
				"Psalm 98",
				"Acts 10:34-48",
				"1 John 5:1-8",
				"John 15:9-17"},
		},
		{"2015-05-14",
			Readings{
				"Psalm 47",
				"Acts 1:1-11",
				"Ephesians 1:15-23",
				"Luke 24:44-53"},
		},
		{"2015-05-17",
			Readings{
				"Psalm 1",
				"Acts 1:12-26",
				"1 John 5:9-15",
				"John 17:11-19"},
		},
		{"2015-05-24",
			Readings{
				"Psalm 139:1-16",
				"Ezekiel 37:1-14",
				"Acts 2:1-21",
				"John 15:26-27, John 16:4-15"},
		},
		{"2015-05-31",
			Readings{
				"Psalm 29",
				"Isaiah 6:1-8",
				"Acts 2:14, Acts 2:22-36",
				"John 3:1-17"},
		},
		{"2015-06-07",
			Readings{
				"Psalm 130",
				"Genesis 3:8-15",
				"2 Corinthians 4:13—5:1",
				"Mark 3:20-35"},
		},
		{"2015-06-14",
			Readings{
				"Psalm 1",
				"Ezekiel 17:22-24",
				"2 Corinthians 5:1-17",
				"Mark 4:26-34"},
		},
		{"2015-06-21",
			Readings{
				"Psalm 124",
				"Job 38:1-11",
				"2 Corinthians 6:1-13",
				"Mark 4:35-41"},
		},
		{"2015-06-28",
			Readings{
				"Psalm 30",
				"Lamentations 3:22-33",
				"2 Corinthians 8:1-9, 2 Corinthians 8:13-15",
				"Mark 5:21-43"},
		},
		{"2015-07-05",
			Readings{
				"Psalm 123",
				"Ezekiel 2:1-5",
				"2 Corinthians 12:1-10",
				"Mark 6:1-13"},
		},
		{"2015-07-12",
			Readings{
				"Psalm 85:1-13",
				"Amos 7:7-15",
				"Ephesians 1:3-14",
				"Mark 6:14-29"},
		},
		{"2015-07-19",
			Readings{
				"Psalm 23",
				"Jeremiah 23:1-6",
				"Ephesians 2:11-22",
				"Mark 6:30-44"},
		},
		{"2015-07-26",
			Readings{
				"Psalm 136:1-9",
				"Genesis 9:8-17",
				"Ephesians 3:14-21",
				"Mark 6:45-56"},
		},
		{"2015-08-02",
			Readings{
				"Psalm 145:10-21",
				"Exodus 16:2-15",
				"Ephesians 4:1-16",
				"John 6:22-35"},
		},
		{"2015-08-09",
			Readings{
				"Psalm 34:1-8",
				"1 Kings 19:1-8",
				"Ephesians 4:17—5:2",
				"John 6:35-51"},
		},
		{"2015-08-16",
			Readings{
				"Psalm 34:12-22",
				"Proverbs 9:1-10, Joshua 24:1-2, Joshua 24:14-18",
				"Ephesians 5:6-21",
				"John 6:51-69"},
		},
		{"2015-08-23",
			Readings{
				"Psalm 14",
				"Isaiah 29:11-19",
				"Ephesians 5:22-33",
				"Mark 7:1-13"},
		},
		{"2015-08-30",
			Readings{
				"Psalm 119:129-136",
				"Deuteronomy 4:1-2, Deuteronomy 4:6-9",
				"Ephesians 6:10-20",
				"Mark 7:14-23"},
		},
		{"2015-09-06",
			Readings{
				"Psalm 146",
				"Isaiah 35:4-7",
				"James 2:1-10, James 2:14-18",
				"Mark 7:24-37"},
		},
		{"2015-09-13",
			Readings{
				"Psalm 116:1-9",
				"Isaiah 50:4-10",
				"James 3:1-12",
				"Mark 9:14-29"},
		},
		{"2015-09-20",
			Readings{
				"Psalm 54",
				"Jeremiah 11:18-20",
				"James 3:13—4:10",
				"Mark 9:30-37"},
		},
		{"2015-09-27",
			Readings{
				"Psalm 104:27-35",
				"Numbers 11:4-6, Numbers 11:10-16, Numbers 11:24-29",
				"James 5:1-20",
				"Mark 9:38-50"},
		},
		{"2015-10-04",
			Readings{
				"Psalm 128",
				"Genesis 2:18-25",
				"Hebrews 2:1-18",
				"Mark 10:2-16"},
		},
		{"2015-10-11",
			Readings{
				"Psalm 90:12-17",
				"Amos 5:6-7, Amos 5:10-15",
				"Hebrews 3:12-19",
				"Mark 10:17-22"},
		},
		{"2015-10-18",
			Readings{
				"Psalm 119:9-16",
				"Ecclesiastes 5:10-20",
				"Hebrews 4:1-16",
				"Mark 10:23-31"},
		},
		{"2015-10-25",
			Readings{
				"Psalm 46",
				"Revelation 14:6-7",
				"Romans 3:19-28",
				"John 8:31-36, Matthew 11:12-19"},
		},
		{"2015-11-01",
			Readings{
				"Psalm 149",
				"Revelation 7:2-17",
				"1 John 3:1-3",
				"Matthew 5:1-12"},
		},
		{"2015-11-08",
			Readings{
				"Psalm 146",
				"1 Kings 17:8-16",
				"Hebrews 9:24-28",
				"Mark 12:38-44"},
		},
		{"2015-11-15",
			Readings{
				"Psalm 16",
				"Daniel 12:1-3",
				"Hebrews 10:11-25",
				"Mark 13:1-13"},
		},
		{"2015-11-22",
			Readings{
				"Psalm 93",
				"Isaiah 51:4-6, Daniel 7:9-10, Daniel 7:13-14",
				"Jude 20-25, Revelation 1:4-8",
				"Mark 13:24-37, John 18:33-37"},
		},
		{"2015-11-26",
			Readings{
				"Psalm 67",
				"Deuteronomy 8:1-10",
				"Philippians 4:6-20, 1 Timothy 2:1-4",
				"Luke 17:11-19"},
		},
		{"2015-11-29",
			Readings{
				"Psalm 25:1-10",
				"Jeremiah 33:14-16",
				"1 Thessalonians 3:9-13",
				"Luke 19:28-40, Luke 21:25-36"},
		},
	}
)

/*type (
	Pericope struct {
		Date     string
		Readings Readings
	}
)*/
/*
type PericopeIterator func() (Pericope, bool)

// Calling the PericopeIterator function returns a Pericope and a
// bool that tells me if the returned value is valid. If it is false,
// I have iterated off the end of the list and the Pericope has some
// useless default value.
func (p Pericope) Iter() PericopeIterator {
    return func(list []Pericope) PericopeIterator {
        index := 0
        return func() (Pericope, bool) {
            if index >= 0 && index < len(list) {
                val := list[index]
                index++
                return val, true
            }
            return Lectionary[0] ,false
        }
    }(p.Date, p.Readings)
}
*/
/*type (
	Readings struct {
		Psalm   []string
		First   []string
		Epistle []string
		Gospel  []string
	}
)*/
/*
type ReadingsIterator func() ([]string, []string, []string, []string, bool)

// Calling the PericopeIterator function returns a Pericope and a
// bool that tells me if the returned value is valid. If it is false,
// I have iterated off the end of the list and the Pericope has some
// useless default value.
func (r Readings) Iter() ReadingsIterator {
    return func(list []Readings) ReadingsIterator {
        index := 0
        return func() (Pericope, bool) {
            if index >= 0 && index < len(list) {
                val := list[index]
                index++
                return val, true
            }
            return Lectionary[0] ,false
        }
    }(r.Psalm, r.First, r.Epistle, r.Gospel)
}
*/
