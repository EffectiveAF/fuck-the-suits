<script>
  import FusionCharts from 'fusioncharts';
  import Charts from 'fusioncharts/fusioncharts.charts';
  import FusionTheme from 'fusioncharts/themes/fusioncharts.theme.candy';
  import SvelteFC, { fcRoot } from './svelte-fusioncharts';
  import dataSource from './data.js';

  export let location = '';

  fcRoot(FusionCharts, Charts, FusionTheme);

  let chartConfigColumn = {
      type: 'column2d',
      renderAt: 'chart-container1',
      width: '100%',
      height: '400',
      dataSource: dataSource.columnData
    },
    chartConfigStackedColumn = {
      type: 'stackedcolumn2d',
      renderAt: 'chart-container2',
      width: '100%',
      height: '400',
      dataSource: dataSource.stackedColumnData
    };

  const exportHandler = () => {
    FusionCharts.batchExport({
      exportFormat: 'pdf'
    });
  };
</script>

<div id='chart-container2'>
  <SvelteFC {...chartConfigStackedColumn} />
</div>
<div style="text-align: left; padding-top: 5px;">
  <button class="btn btn-outline-secondary btn-sm" on:click={exportHandler}>Export Both charts as a single pdf</button>
</div>
