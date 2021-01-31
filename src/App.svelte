<script>
  import { Router, Route } from 'svelte-routing';

  import Home from './Home.svelte';
  import StackedBarChart from './StackedBarChart.svelte';
  import TermsOfService from './TermsOfService.svelte';
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
    <Route path="/tos" component="{TermsOfService}" />
    <Route path="/chart/2"><StackedBarChart /></Route>
    <Route path="/"><Home /></Route>
  </div>
</Router>
