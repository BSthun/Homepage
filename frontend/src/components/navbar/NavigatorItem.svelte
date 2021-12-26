<script lang="ts">
	import { navigate } from 'svelte-navigator';
	
	export let la: string;
	export let href: string;
	export let title: string;
	export let active: boolean;
	export let exact: boolean;
	export let order: number;
	export let toggled: boolean;
	
	let w = 0;
	let ww = 0;
	
	const nav = () => {
		navigate(href, { replace: exact });
	};
</script>

<svelte:window bind:innerWidth={ww} />

<div class="link"
     on:click={() => nav()}
     style={ww < 768 ? `opacity: ${toggled ? 1 : 0}; transform: translateY(${toggled ? 0 : 12}px); transition: opacity .5s ease-in-out ${toggled ? order * 0.3: 1.2 - order * 0.3}s` : ''}
>
	<i class="{la}"></i>
	<div class="box"
	     style={`width: ${ww < 768 || active ? w + 4 : 0}px`}
	>
		<p bind:offsetWidth={w}>{title}</p>
	</div>
</div>

<style lang="scss">
	.link {
		display: flex;
		align-items: center;
		margin-right: 16px;
		color: $color-grey-100;
		opacity: 0.8;
		
		&:hover {
			opacity: 1;
		}
		
		@include breakpoint('md', 'dn') {
			margin: 0;
			padding: 16px 8px;
		}
		
		&:not(:last-child) {
			@include breakpoint('md', 'dn') {
				border-bottom: 1px solid $color-white;
			}
		}
		
		i {
			font-size: 24px;
		}
		
		.box {
			width: 0;
			overflow: hidden;
			transition: width 0.3s ease-in-out;
			
			p {
				margin-left: 4px;
				width: fit-content;
			}
		}
	}
</style>