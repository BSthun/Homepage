<script lang="ts">
	import { trackLog } from '../../utils/api/track'
	import { useLocation } from 'svelte-navigator'
	import ImgBSthunFlat from '../../images/logo/bsthun-flat-white.svg'
	import NavigatorItem from './NavigatorItem.svelte'
	import type { NavbarItem } from './types'

	const location = useLocation()

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
			title: 'Photo',
			la: 'las la-camera',
			href: '/photo',
		},
	]

	let toggled = false
	let toggle = (log = true) => {
		if (window.innerWidth < 768) {
			toggled = !toggled
		} else {
			toggled = false
		}

		if (log) {
			trackLog('navbar/toggle', null, toggled)
		}
	}
</script>

<nav class="navbar">
	<div>
		<div class="title">
			<img alt="BSthun Flat Logo" src={ImgBSthunFlat} />
			<h1>BSthun</h1>
		</div>
		<div class={`navigator ${toggled && 'scaled'}`}>
			{#each items as item, i}
				<NavigatorItem
					active={item.href === '/' ? $location.pathname === '/' : $location.pathname.startsWith(item.href)}
					exact={item.href === $location.pathname}
					title={item.title}
					href={item.href}
					la={item.la}
					order={i + 1}
					{toggled}
					{toggle}
				/>
			{/each}
		</div>

		<div class="hamburger" on:click={toggle} on:keypress={toggle}>
			<i class="las la-grip-lines" />
		</div>

		<div class={`circular ${toggled && 'scaled'}`} />
	</div>
</nav>

<style lang="scss">
	@import 'src/styles/_index.scss';

	nav {
		position: fixed;
		top: 0;
		height: $size-navbar-height;
		width: 100%;
		color: $color-grey-200;
		background: $color-navbar-bg-scrolled;
		backdrop-filter: blur(4px);
		-webkit-backdrop-filter: blur(4px);
		border-bottom: 1px solid $color-border;
		z-index: 1201;

		> div {
			position: relative;
			width: 100%;
			height: 100%;
			max-width: map-get($breakpoints, 'xl');
			margin: auto;
			display: flex;
			align-items: center;
			justify-content: space-between;

			.title {
				display: flex;
				margin-left: 16px;
				z-index: 1204;

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
				z-index: 1203;

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
				z-index: 1205;
				left: 16px;
				font-size: 24px;
				cursor: pointer;

				@include breakpoint('md', 'up') {
					display: none;
				}
			}

			.circular {
				position: absolute;
				z-index: 1200;
				border-radius: $size-infinity;
				left: #{24px - 128px};
				top: #{$size-navbar-height * 0.5 - 128px};
				width: 256px;
				height: 256px;
				background-color: $color-bluegrey-700;

				transform: scale3d(0, 0, 1);
				transform-origin: center;
				transition-property: all;
				transition-duration: 0.5s;
				transition-timing-function: $animate-easing-nav;
				transition-delay: 0.2s;

				&.scaled {
					transform: scale3d(10, 10, 1);
					transition-delay: 0s;
				}
			}
		}
	}
</style>
