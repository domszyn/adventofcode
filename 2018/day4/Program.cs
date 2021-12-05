using System;
using System.Text.RegularExpressions;

var schedule = new Dictionary<int, Dictionary<DateTime, Dictionary<int, bool>>>();

var guardRegex = new Regex(@"#(?<id>\d+)");

var guardID = 0;
DateTime date = new DateTime();
foreach (var record in Input.Records)
{
    if (record.EndsWith("begins shift"))
    {
        guardID = Convert.ToInt32(guardRegex.Matches(record)[0].Groups["id"].Value);
        if (!schedule.ContainsKey(guardID))
        {
            schedule[guardID] = new Dictionary<DateTime, Dictionary<int, bool>>();
        }
    }
    date = DateTime.Parse(record.Substring(1, 16));
    var dateKey = date - date.TimeOfDay;
    if (date.Hour != 0)
    {
        dateKey += TimeSpan.FromDays(1);
    }

    if (!schedule[guardID].ContainsKey(dateKey))
    {
        schedule[guardID][dateKey] = new Dictionary<int, bool>();
        for (var i = 0; i < 60; i++)
        {
            schedule[guardID][dateKey][i] = false;
        }
    }

    if (record.EndsWith("falls asleep"))
    {
        schedule[guardID][dateKey][date.Minute] = true;
    }

    if (record.EndsWith("wakes up"))
    {
        for (var i = date.Minute - 1; i >= 0; i--)
        {
            if (schedule[guardID][dateKey][i]) break;

            schedule[guardID][dateKey][i] = true;
        }
    }
}

var maxMinutesAsleep = 0;
var maxAsleepID = 0;
foreach (var id in schedule.Keys)
{
    var minutesAsleep = 0;
    foreach (var day in schedule[id].Values)
    {
        minutesAsleep += day.Count(_ => _.Value);
    }

    if (minutesAsleep > maxMinutesAsleep)
    {
        maxMinutesAsleep = minutesAsleep;
        maxAsleepID = id;
    }
}

var maxDaysAsleep = 0;
var maxMinuteAsleep = -1;
for (var i = 0; i < 60; i++)
{
    var daysAsleep = 0;
    foreach (var day in schedule[maxAsleepID].Values)
    {
        if (day[i])
        {
            daysAsleep++;
        }
    }
    if (daysAsleep > maxDaysAsleep)
    {
        maxDaysAsleep = daysAsleep;
        maxMinuteAsleep = i;
    }
}

Console.WriteLine("Part 1: {0}", maxMinuteAsleep * maxAsleepID);

var mostFrequentID = 0;
var mostFrequentMinute = 0;
var mostDaysAsleep = 0;

for (var i = 0; i < 60; i++)
{
    foreach (var id in schedule.Keys)
    {
        var daysAsleep = 0;
        foreach (var day in schedule[id].Values)
        {
            if (day[i])
            {
                daysAsleep++;
            }
        }
        if (daysAsleep > mostDaysAsleep)
        {
            mostDaysAsleep = daysAsleep;
            mostFrequentID = id;
            mostFrequentMinute = i;
        }
    }
}

Console.WriteLine("Part 2: {0}", mostFrequentID*mostFrequentMinute);