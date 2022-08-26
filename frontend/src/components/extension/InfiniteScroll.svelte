<script>
	import {
		createEventDispatcher,
		onDestroy,
	} from 'svelte';
	
	export let threshold = 0;
	export let more = true;
	
	const dispatch = createEventDispatcher();
	let isLoadMore = false;
	let component;
	
	const onScroll = (e) => {
		if (document.documentElement.scrollHeight - window.innerHeight - window.scrollY <= threshold) {
			if (!isLoadMore && more) {
				dispatch('load');
			}
			isLoadMore = true;
		} else {
			isLoadMore = false;
		}
	};
	
	$: {
		if (component) {
			window.addEventListener('scroll', onScroll);
			window.addEventListener('resize', onScroll);
		}
	}
	
	onDestroy(() => {
		window.addEventListener('scroll', onScroll);
		window.addEventListener('resize', onScroll);
	});
</script>

<div bind:this={component} style="width:0"></div>