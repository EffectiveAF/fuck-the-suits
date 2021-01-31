<script>
  import Footer from './Footer.svelte';

  export let location = '';
</script>

<div class="home container-fluid">
  <div class="section-one">
    <h2>Ready to change the World?</h2>
  </div>

  <div class="section-one" style="margin-top: 100px">
    <h2>Get started today.</h2>
    <p class="p-large p-large-ctn" style="text-align: center; padding: 24px 0;">We imagine a world where anyone, anywhere can create a highly productive, volunteer-led, 100-person organization — in 1 hour.</p>
    <button class="button primary">Sign up, cuck</button>
  </div>

  <Footer />

</div>

<style>
  .home {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .section-one {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 85%;
    text-align: center;
    padding: 100px 0;
  }

  .p-large-ctn {
    width: 80%;
  }

  @media (max-width: 499px) {
    .p-large {
      font-size: 18px;
    }

    .p-large-ctn {
      width: 96%;
    }

    .section-one {
      width: 100%;
    }
  }
</style>
