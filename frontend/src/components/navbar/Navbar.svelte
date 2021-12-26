<script lang="ts">
	import ImgBSthunFlat from '../../images/logo/bsthun-flat-white.svg';
	import { useLocation } from 'svelte-navigator';
	import type { NavbarItem } from './types';
	import NavigatorItem from './NavigatorItem.svelte';
	
	const location = useLocation();
	
	const items: NavbarItem[] = [
		{
			context: null,
			title: 'Home',
			la: 'las la-home',
			href: '/',
		},
		{
			context: null,
			title: 'Project',
			la: 'las la-lightbulb',
			href: '/project',
		},
		{
			context: null,
			title: 'Blog',
			la: 'las la-pen-nib',
			href: '/blog',
		},
		{
			context: null,
			title: 'Photograph',
			la: 'las la-camera',
			href: '/photograph',
		},
	];
	
	let toggled = false;
</script>

<nav>
	<div>
		<div class="title">
			<img alt="BSthun Flat Logo" src={ImgBSthunFlat}>
			<h1>BSthun</h1>
		</div>
		<div class={`navigator ${toggled && 'scaled'}`}>
			{#each items as item, i}
				<NavigatorItem active={item.href === '/' ? $location.pathname === '/': $location.pathname.startsWith(item.href)}
				               exact={item.href === $location.pathname}
				               title={item.title}
				               href={item.href}
				               la={item.la}
				               order={i + 1}
				               toggled={toggled}
				/>
			{/each}
		</div>
		
		<div class="hamburger" on:click={() => toggled = !toggled}>
			<i class="las la-grip-lines"></i>
		</div>
		
		<div class={`circular ${toggled && 'scaled'}`}></div>
	</div>
</nav>

<style lang="scss">
	nav {
		height: $size-navbar-height;
		width: 100%;
		color: $color-grey-100;
		background-color: $color-bluegrey-900;
		
		> div {
			position: relative;
			width: 100%;
			height: 100%;
			max-width: map-get($breakpoints, 'lg');
			margin: auto;
			display: flex;
			align-items: center;
			justify-content: space-between;
			
			.title {
				display: flex;
				margin-left: 16px;
				z-index: 124;
				
				@include breakpoint('md', 'dn') {
					margin-left: 52px;
				}
				
				img {
					width: 32px;
				}
				
				h1 {
					margin-left: 8px;
					letter-spacing: 1px;
					line-height: 1.33;
				}
			}
			
			.navigator {
				display: flex;
				z-index: 123;
				
				&:not(.scaled) {
					@include breakpoint('md', 'dn') {
						transition: visibility 1s step-end;
						visibility: hidden;
					}
				}
				
				@include breakpoint('md', 'dn') {
					position: absolute;
					top: 0;
					left: 0;
					right: 0;
					height: 100vh;
					flex-direction: column;
					justify-content: center;
					width: 200px;
					margin: auto;
					transition: none;
				}
			}
			
			.hamburger {
				position: absolute;
				z-index: 125;
				left: 16px;
				font-size: 24px;
				cursor: pointer;
				
				@include breakpoint('md', 'up') {
					display: none;
				}
			}
			
			.circular {
				position: absolute;
				z-index: 9;
				border-radius: $size-infinity;
				left: #{24px - 128px};
				top: #{$size-navbar-height / 2 - 128px};
				transform-origin: center;
				width: 256px;
				height: 256px;
				transform: scale3d(0, 0, 1);
				transition: all 0.7s ease-in-out 1s;
				
				background-color: $color-bluegrey-800;
				
				&.scaled {
					transform: scale3d(10, 10, 1);
					transition: all 0.7s ease-in-out;
				}
			}
		}
	}
</style>