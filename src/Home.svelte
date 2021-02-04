<script>
  import { navigate } from 'svelte-routing';
  import FusionCharts from 'fusioncharts';
  import Charts from 'fusioncharts/fusioncharts.charts';
  import FusionTheme from 'fusioncharts/themes/fusioncharts.theme.candy';
  import SvelteFC, { fcRoot } from './svelte-fusioncharts';
  // import StackedBarChart from './StackedBarChart.svelte';
  import BarChart from './BarChart.svelte';
  import TimeSeriesChart from './TimeSeriesChart.svelte';
  import { reactivePath, currentSymbol, chartData } from './stores.js';

  export let location = '';
  export let symbol = '';

  if (symbol) {
    $currentSymbol = symbol;
  }
</script>

<!-- <div class="container" style="border: 1px solid red;"> -->
<div class="container">

  <div class="ex-side">
    <div class="sidebar-content">
      <div class="logo-ctn">
        <img src="/img/fuck-the-suits.svg" alt="In the end, We the People will win." height="65px" width="103px">
        <div class="eff-the-suits"></div>
      </div>

      <div class="sidebar-items">
        <p class="labels-overline">Top Shorts<br />In Market</p>
        <ul>
          <li>
            <button class={"link " + ($reactivePath === '/' ? 'active-link' : '')} on:click={(e) => {
              e.preventDefault();
              navigate('/');
              $currentSymbol = '';
            }}>
            The Biggest Shorts
            </button>
          </li>

          {#each $chartData as company (company.label)}
          <li>
            <button class={"link " + ($reactivePath === '/chart/' + company.label ? 'active-link' : '')} on:click={(e) => {
              e.preventDefault();
              navigate('/chart/' + company.label);
              $currentSymbol = company.label;
            }}>
              {company.label}
            </button>
          </li>
          {/each}
        </ul>

        <div style="height: 32px;"></div>

        <p class="labels-overline">More Info</p>
        <ul>
          <li>
            <button class="link" on:click={(e) => {
              e.preventDefault();
              navigate('/methodology');
              $currentSymbol = '';
            }}>
              Methodology
            </button>
          </li>
        </ul>
      </div>
    </div>
  </div>
  <div class="main-grid">
    {#if !$currentSymbol}
      <div class="section">
        <div class="header-and-image">
          <h1>The Biggest Shorts</h1>
          <div style="width: 16px;"></div>
          <img src="/img/short-shorts2.jpg" alt="Who wears short shorts?" />
        </div>

        <p class="p-large" style="padding-top: 8px;">
          Below are the most shorted companies in the world -- those with the
          highest short volume -- over the last 5 trading days.
          <br />
          <br />
          You know what to do. #HoldTheLine
          <br />
          <br />
        </p>
      </div>

      <BarChart />
    {:else}
      <TimeSeriesChart />
    {/if}
  </div>
</div>

<style>
.header-and-image {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.section {
  /* border: 1px solid red; */
  height: auto;
  width: 540px;
}

/* Sidebar */
.sidebar-items {
  display: flex;
  flex-direction: column;
  width: calc(100% - 32px);
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
  height: 1000px;
  flex-direction: row;
  background-color: #181818;
  border-radius: 8px;
  box-shadow: var(--pop);
  margin: 32px 0;
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

.eff-the-suits {
  background-image: url('/img/wsb_dude.png');
  width: 65px;
  height: 65px;
  margin-left: 8px;
}

.eff-the-suits:hover {
  cursor: pointer;
  background-image: url('/img/anon.png');
}

.link {
  color: var(--contentColorDefault);
  text-decoration: none;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 20px;
  letter-spacing: 0.0025em;
  margin: 6px 0;
  padding: 1px;
  border-left: 1px solid white;
  border-radius: 0px;
  padding-left: 4px;
}

.active-link,
.link:hover {
  background-color: #fff;
  color: #2c2c2c;
  border-radius: 0px;
  -webkit-transform: translateX(5px); /* 10px to left */
  -moz-transform: translateX(5px); /* 10px to left */
  -ms-transform: translateX(5px); /* 10px to left */
  transform: translateX(5px); /* 10px to left */
  transition: .05s ease-in-out;
}

.link:focus {
  -webkit-focus-ring-color: auto 5px;
}
</style>
