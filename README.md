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

...assuming the top order didn't change:

```sql
fuckthesuits=# do $body$ declare _symbol text; _date text; _short_volume int; symbols text[] := array['SNDL', 'NAKD', 'CTRM', 'AMC', 'NOK', 'ZOM']; begin FOREACH _symbol IN ARRAY symbols LOOP SELECT date, short_volume FROM daily_short_volume WHERE symbol = _symbol AND date = '2021-02-03' INTO _date, _short_volume; RAISE NOTICE '%: ["%", %]', _symbol, _date, _short_volume; END LOOP; end; $body$;
```


## Attributions

- Priyanjit Dey for our fork of `svelte-fusioncharts/`
