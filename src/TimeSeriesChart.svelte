<script>
  import FusionCharts from 'fusioncharts';
  import Timeseries from 'fusioncharts/fusioncharts.timeseries';
  import SvelteFC, { fcRoot } from './svelte-fusioncharts';

  import { currentSymbol, companyTimeSeriesData } from './stores.js';

  const timeSeriesSchema = [
    {
      "name": "Date",
      "type": "date",
      "format": "%Y-%m-%d"
    },
    {
      "name": "Short Volume",
      "type": "number"
    }
  ];

  fcRoot(FusionCharts, Timeseries);

  const getChartConfig = ([data, schema]) => {
    const fusionDataStore = new FusionCharts.DataStore();
    const fusionTable = fusionDataStore.createDataTable(data, schema);

    return {
      type: 'timeseries',
      width: '100%',
      height: 450,
      renderAt: 'chart-container',
      dataSource: {
        data: fusionTable,
        caption: {
          text: 'New shorts against ' + $currentSymbol + ' (none on weekends)'
        },
        subcaption: {
          text: ''
        },
        yAxis: [
          {
            plot: {
              value: 'Short Volume',
              type: 'column'
            },
            format: {
              prefix: ''
            },
            title: 'Short Volume'
          }
        ]
      }
    };
  };
</script>

{#if $currentSymbol}
<div id="chart-container">
  <SvelteFC
    {...getChartConfig([$companyTimeSeriesData[$currentSymbol], timeSeriesSchema])}
  />
</div>
{/if}
