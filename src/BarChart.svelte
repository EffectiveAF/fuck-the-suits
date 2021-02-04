<script>
  import FusionCharts from 'fusioncharts';
  import Charts from 'fusioncharts/fusioncharts.charts';
  import FusionTheme from 'fusioncharts/themes/fusioncharts.theme.candy';
  import SvelteFC, { fcRoot } from './svelte-fusioncharts';
  import { currentSymbol, chartData } from './stores.js';

  // Always set FusionCharts as the first parameter
  fcRoot(FusionCharts, Charts, FusionTheme);

  const dataSource = {
    chart: {
      caption: 'The Most Shorted Companies This Past Week',
      subCaption: '',
      xAxisName: 'Company Ticker Symbol',
      yAxisName: 'Short Volume (last 5 trading days)',
      numberSuffix: '',
      theme: 'candy'
    },
    data: $chartData
  };

  const chartConfigs = {
    type: 'column2d',
    width: '100%',
    height: 400,
    dataFormat: 'json',
    dataSource: dataSource
  };
</script>

{#if $currentSymbol === ''}
<div class="data-viz">
  <SvelteFC {...chartConfigs} />
</div>
{/if}

<style>
.data-viz {
  padding-top: 8px;
}

</style>
