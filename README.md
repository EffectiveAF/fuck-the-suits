# Fuck the Suits

For far too long, The Suits have served their own selfish interests
above the rights of the majority -- and gotten away with it.

Until now.


## Run Dev Environment

```
# Clone this repo
npm i
npm run build
npm run dev

# Or
npm i && npm run build && npm run dev
```

Now visit <http://localhost:5000/> !


## Getting short_volume Data

Top 6 most shorted companies in the last 5 trading days, assuming
there were no holidays in the last week:

```
fuckthesuits=# SELECT symbol, SUM(short_volume) AS total FROM daily_short_volume WHERE date > now() - interval '1d' * 7 GROUP BY symbol ORDER BY total DESC LIMIT 6;
```

Get the latest data on the above companies, assuming the top 6 didn't change:

```sql
fuckthesuits=# do $body$ declare _symbol text; _date text; _short_volume int; symbols text[] := array['SNDL', 'NAKD', 'AMC', 'CTRM', 'ZOM', 'NOK']; begin FOREACH _symbol IN ARRAY symbols LOOP SELECT date, short_volume FROM daily_short_volume WHERE symbol = _symbol AND date = '2021-02-03' INTO _date, _short_volume; RAISE NOTICE '%: ["%", %]', _symbol, _date, _short_volume; END LOOP; end; $body$;
```


## Attributions

- Priyanjit Dey for our fork of `svelte-fusioncharts/`
