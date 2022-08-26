<script lang="ts">
	import { getContext } from 'svelte';
	import Zoom from 'svelte-zoom';
	
	const gallery = getContext('gallery');
	$: item = $gallery.items[$gallery.expand];
	
	$: (() => {
		if ($gallery.expand != null)
			document.querySelector('body').style.overflow = 'hidden';
		else
			document.querySelector('body').style.overflow = 'scroll';
		
	})();
</script>

{#if ($gallery.expand != null)}
	<div class="gallery-detail">
		<div class="img-wrapper">
			<Zoom alt={item.title} src={item.root + item.image_path} />
		</div>
		<div class="plane">
		
		</div>
	</div>
{/if}

<style lang="scss">
	@import 'src/styles/_index.scss';
	
	.gallery-detail {
		position: fixed;
		inset: 0;
		z-index: 10000;
		background-color: #4E809820;
		backdrop-filter: blur(4px) brightness(0.5);
		-webkit-backdrop-filter: blur(4px) brightness(0.5);
		display: flex;
		
		@include breakpoint('md', 'dn') {
			flex-direction: column;
		}
	}
	
	.img-wrapper {
		position: relative;
		flex: 1;
		overflow: hidden;
		border-radius: 8px;
		margin: 12px;
	}
	
	.plane {
		display: flex;
		background-color: white;
		width: 300px;
		border-radius: 8px;
		margin: 12px;
		
		@include breakpoint('md', 'dn') {
			width: 100%;
		}
	}
</style>