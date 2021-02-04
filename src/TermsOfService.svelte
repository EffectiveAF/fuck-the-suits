<script>
  import Footer from './Footer.svelte';

  export let location = '';
</script>

<div class="container">
  <div class="main-grid">

  <h2>Terms of Service</h2>

  <h4>Guiding philosophy: you can do anything you want with this software <em>except</em> commit or plan human rights violations.<strong></strong></h4>

  <p>See below for details and official legalese version.</p>

  <p>(TODO: Add remainder of TOS here before public launch.)</p>

  </div>
</div>

<style>
p {
  font-size: 18px;
}

h2, h4, p {
  margin-top: 24px;
  max-width: 80%;
}

@media (max-width: 639px) {
  h2, h4, p {
    max-width: calc(100% - 24px);
  }
}
</style>
