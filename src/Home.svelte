<script>
  import Footer from './Footer.svelte';
  import FusionCharts from 'fusioncharts';
  import Charts from 'fusioncharts/fusioncharts.charts';
  import FusionTheme from 'fusioncharts/themes/fusioncharts.theme.candy';
  import SvelteFC, { fcRoot } from './svelte-fusioncharts';
  import Divider from './Divider.svelte';
  import StackedBarChart from './StackedBarChart.svelte';

  import { chartData } from './stores.js';

  export let location = '';

  // Always set FusionCharts as the first parameter
  fcRoot(FusionCharts, Charts, FusionTheme);

  const dataSource = {
    chart: {
      caption: 'It is time to short stock the short stockers, motherfuckers.',
      subCaption: 'In MMbbl = One Million barrels',
      xAxisName: 'Country',
      yAxisName: 'Reserves (MMbbl)',
      numberSuffix: 'K',
      theme: 'candy'
    },
    data: $chartData
  };

  const chartConfigs = {
    type: 'column2d',
    width: "100%",
    height: 400,
    dataFormat: 'json',
    dataSource: dataSource
  };
</script>

<!-- <div class="container" style="border: 1px solid red;"> -->
<div class="container">
  
  <div class="ex-side">
    <Divider />
    <div class="sidebar-content">
      <div class="logo-ctn">
        <img src="/img/fuck-the-suits.svg" alt="The People will win, always." height="65px" width="103px">
        <img src="/img/lul.gif" alt="We Do Not Forgive. We Do Not Forget." height="65px" width="65px" style="padding-left: 8px">
      </div>

      <div class="sidebar-items">
        <p class="labels-overline">Data Visualization</p>
        <ul>
          <li>List Item 1</li>
          <li>List Item 2</li>
          <li>List Item 3</li>
          <li>List Item 4</li>
        </ul>
      </div>
    </div>
    <Divider />
  </div>
  <div class="main-grid">
    <div class="section">
      <h1>Header</h1>
      <p class="p-large" style="padding-top: 8px">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Auctor faucibus eget mauris ornare ornare. Mauris, molestie. Change the world.</p>
    </div>
    <div class="data-viz">
      <SvelteFC {...chartConfigs} />
      <div style="padding-top: 16px;"></div>
      <StackedBarChart />
    </div>
  </div>
 
</div>

<Footer />

<style>
.container {
  display: flex;
  height: 100%;
  min-height: 100vh;
  max-width: 1250px;
  margin: 0 auto;
}

.main-grid {
  max-width: 840px;
  width: 100%;
  height: 100%;
  margin: 0 auto;
  padding-top: 176px;
  /* border: 1px solid yellow; */
}

.section {
  /* border: 1px solid red; */
  height: auto;
  width: 488px;
}

/* Sidebar */
.sidebar-items {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: auto;
  padding: 80px 16px 0px 16px;
}

.sidebar-items ul {
  padding-top: 8px;
}

.sidebar-items ul li {
  display: flex;
  flex-direction: row;
  font-family: "IBM Plex Sans";
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: 30px;
  letter-spacing: 0.0025em;
}

.ex-side {
  display: flex;
  width: 210px;
  height: 100%;
  flex-direction: row;
}

.sidebar-content {
  display: flex;
  flex-direction: column;
  width: calc(100% - 2px);
  /* border: 1px solid yellow; */
  -ms-overflow-style: none;
}

.ex-side a {
  font-size: 0.9em;
}

.ex-side-content {
  position: fixed;
  width: 185px;
  overflow-y: auto;
  overflow-x: hidden;
  height: calc(100% - 10vh);
  -ms-overflow-style: none;
}

.ex-side-content::-webkit-scrollbar {
  width: 0 !important;
  display: none !important;
}

.ex-side-content li {
  list-style: none;
}

.ex-side-content li:before {
  content: '';
}

.ex-side-content ul {
  margin-left: 0;
}

.zi-title {
  color: var(--contentColorActive);
  text-transform: uppercase;
  font-size: 0.8em;
}

.logo-ctn {
  display: flex;
  flex-direction: row;
  padding: 56px 8px 0px 16px;
  height: 64px;
}

/* Data Viz */
.data-viz {
  padding-top: 8px;
}

:global(#raphael-paper-2) {
  border: 1px solid #818285;
}

</style>
