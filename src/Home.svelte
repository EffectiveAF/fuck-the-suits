<script>
  import Footer from './Footer.svelte';

  export let location = '';
</script>

<div class="container" style="border: 1px solid red;">
  
  <div class="ex-side">
    <div class="divider"></div>
    <div class="sidebar-content">
      <div class="logo-ctn">
        <img src="/img/fuck-the-suits.svg" alt="The People will win, always." height="65px" width="103px">
        <img src="/img/lul.gif" alt="We Do Not Forgive. We Do Not Forget." height="65px" width="65px" style="padding-left: 8px">
      </div>
    </div>
    <div class="divider"></div>
  </div>
  <div class="main-grid">

  </div>

</div>

<Footer />

<style>
.container {
  display: flex;
  min-height: 100vh;
  max-width: 1250px;
  margin: 0 auto;
}

.main-grid {
  max-width: 840px;
  width: 100%;
  height: 100vh;
  margin: auto;
  border: 1px solid yellow;
}

/* Sidebar */
.ex-side {
  display: flex;
  width: 200px;
  height: 100vh;
  flex-direction: row;
}

.sidebar-content {
  display: flex;
  flex-direction: row;
  width: 198px;
  border: 1px solid yellow;
}

.divider {
  display: flex;
  flex-direction: row;
  width: 1px;
  height: 100%;
  background-color: white;
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
</style>
