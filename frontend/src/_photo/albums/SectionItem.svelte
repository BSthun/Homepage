<script lang="ts">
	import { navigate } from 'svelte-navigator';
	import moment from 'moment';
	import SectionMeta from './SectionMeta.svelte';
	
	export let item: object;
	
	$: console.log(item);
	
	const onClick = () => {
		navigate('/photo/section/' + item.id);
	};
</script>

<div class="section-item" on:click={() => onClick()}>
	<div
		class="image"
		style="background-image: url('https://www.bsthun.com/photograph/photograph_1_thumbnail.jpg')"
	></div>
	<SectionMeta item={item}>
	<div class="overlay"></div>
	<div class="image-overlay"></div>
</div>

<style lang="scss">
	.section-item {
		width: 100%;
		height: 420px;
		border-radius: 20px;
		overflow: hidden;
		position: relative;
		cursor: pointer;
		
		&:hover {
			.image-overlay {
				backdrop-filter: blur(4px) brightness(0.9);
			}
		}
	}
	
	.image {
		position: absolute;
		inset: 0;
		z-index: 0;
	}
	
	.info {
		position: absolute;
		bottom: 0;
		display: flex;
		gap: 4px;
		flex-direction: column;
		z-index: 10;
		padding: 24px;
		
		.info-meta {
			display: flex;
			align-items: center;
			margin-top: 4px;
			
			i {
				font-size: 22px;
				padding-right: 6px;
			}
			
			h4 {
				font-weight: 400;
				padding-right: 12px;
			}
		}
	}
	
	.overlay {
		position: absolute;
		background: linear-gradient(0deg, rgba(0, 0, 0, 0.7) 0%, rgba(0, 0, 0, 0) 100%);
		bottom: 0;
		left: 0;
		right: 0;
		height: 200px;
		z-index: 5;
	}
	
	.image-overlay {
		position: absolute;
		inset: 0;
		z-index: 1;
		transition: backdrop-filter 300ms ease-in-out;
	}
</style>