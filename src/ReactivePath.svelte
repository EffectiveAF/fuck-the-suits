<script>
  import { onMount, onDestroy } from 'svelte';
  import { reactivePath } from './stores.js';
  import { globalHistory } from 'svelte-routing/src/history.js';

  const updateReactivePath = () => {
    $reactivePath = document.location.pathname;
  };

  let unlisten;
  onMount(() => {
    unlisten = globalHistory.listen(_history => {
      $reactivePath = _history.location.pathname;
    });
  });

  onDestroy(() => {
    unlisten();
  });
</script>


<svelte:window
  on:popstate={(e) => {
    updateReactivePath();
  }}
  on:pushstate={(e) => {
    updateReactivePath();
  }}
/>
