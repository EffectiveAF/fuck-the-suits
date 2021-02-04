<script>
  import { Router, Route } from 'svelte-routing';

  import Home from './Home.svelte';
  import StackedBarChart from './StackedBarChart.svelte';
  import Methodology from './Methodology.svelte';
  import TermsOfService from './TermsOfService.svelte';
  import Footer from './Footer.svelte';
  import FeedbackPopover from './FeedbackPopover.svelte';
</script>

<style>
  nav {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
    border-right: 1px solid var(--neutral200);
    background-color: black;
    width: 320px;
  }
</style>

<Router>
  <nav>
  </nav>
  <div>
    <Route path="/chart/2" component="{StackedBarChart}" />
    <Route path="/chart/:symbol" component="{Home}" />
    <Route path="/methodology" component="{Methodology}" />
    <Route path="/tos" component="{TermsOfService}" />
    <Route path="/"><Home /></Route>
  </div>
  <Footer />
</Router>

<FeedbackPopover />
