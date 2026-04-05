<template>
  <div class="search-page">
    <div class="cyber-bg">
      <div class="grid-overlay"></div>
      <div class="scan-line"></div>
      <div class="floating-particles">
        <div class="particle" v-for="n in 20" :key="n" :style="particleStyle(n)"></div>
      </div>
    </div>

    <div class="search-header">
      <div class="logo-area">
        <router-link to="/" class="logo-link">
          <div class="cyber-logo">
            <span class="logo-text">HIX</span><span class="logo-text-accent">NAV</span>
          </div>
        </router-link>
      </div>
      <div class="search-box">
        <div class="search-input-wrapper">
          <div class="input-glow"></div>
          <input type="text" v-model="keyword" @keyup.enter="doSearch" @input="onInput" placeholder="搜索导航、链接..." class="cyber-input" ref="searchInput" />
          <button class="cyber-btn" @click="doSearch">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8"/><path d="m21 21-4.35-4.35"/>
            </svg>
            <span>搜索</span>
          </button>
        </div>
        <div class="search-suggestions" v-if="showSuggestions && suggestions.length > 0">
          <div v-for="(item, idx) in suggestions" :key="idx" class="suggestion-item" @click="selectSuggestion(item)">
            <img :src="item.logo || defaultLogo" alt="" class="sug-logo" />
            <div class="sug-info">
              <span class="sug-name" v-html="highlightText(item.name)"></span>
              <span class="sug-cat">{{ item.catename }}</span>
            </div>
            <span class="sug-type" :class="'type-' + item.type">{{ item.type === 'nav' ? '导航' : '链接' }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="search-results-container">
      <div class="loading-state" v-if="loading">
        <div class="cyber-loader">
          <div class="loader-ring"></div>
          <div class="loader-ring"></div>
          <div class="loader-ring"></div>
          <span class="loader-text">SEARCHING...</span>
        </div>
      </div>

      <div class="empty-state" v-else-if="searched && navResults.length === 0 && linkResults.length === 0">
        <div class="empty-icon">
          <svg viewBox="0 0 24 24" width="64" height="64" fill="none" stroke="currentColor" stroke-width="1.5">
            <circle cx="11" cy="11" r="8"/><path d="m21 21-4.35-4.35"/><path d="M8 11h6"/>
          </svg>
        </div>
        <p class="empty-text">未找到与 "<span>{{ keyword }}</span>" 相关的结果</p>
        <p class="empty-hint">试试其他关键词吧</p>
      </div>

      <div class="results-content" v-else-if="navResults.length > 0 || linkResults.length > 0">
        <div class="results-stats">
          <span class="stat-text">找到约 <em>{{ totalResults }}</em> 条结果</span>
          <span class="stat-time">(用时 {{ searchTime }} 秒)</span>
        </div>

        <div class="result-section" v-if="navResults.length > 0">
          <div class="section-header">
            <div class="header-line"></div>
            <h2 class="section-title">
              <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/></svg>
              导航结果
              <span class="count-badge">{{ navResults.length }}</span>
            </h2>
            <div class="header-line"></div>
          </div>
          <div class="result-grid">
            <div class="result-card cyber-card" v-for="item in navResults" :key="'nav-' + item.id" @click="jumpTo(item.url)">
              <div class="card-corner corner-tl"></div>
              <div class="card-corner corner-tr"></div>
              <div class="card-corner corner-bl"></div>
              <div class="card-corner corner-br"></div>
              <div class="card-scan"></div>
              <div class="card-content">
                <div class="card-header">
                  <img :src="item.logo || defaultLogo" alt="" class="card-logo" />
                  <div class="card-title-area">
                    <h3 class="card-title" v-html="highlightText(item.name)"></h3>
                    <span class="card-cat">{{ item.catename }}</span>
                  </div>
                  <span class="card-type type-nav">导航</span>
                </div>
                <p class="card-desc" v-html="highlightText(item.desc)"></p>
                <div class="card-footer">
                  <span class="card-url">{{ item.url }}</span>
                  <svg class="card-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M5 12h14M12 5l7 7-7 7"/>
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="result-section" v-if="linkResults.length > 0">
          <div class="section-header">
            <div class="header-line"></div>
            <h2 class="section-title">
              <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></svg>
              快链结果
              <span class="count-badge">{{ linkResults.length }}</span>
            </h2>
            <div class="header-line"></div>
          </div>
          <div class="result-grid">
            <div class="result-card cyber-card link-card" v-for="item in linkResults" :key="'link-' + item.id" @click="jumpTo(item.url)">
              <div class="card-corner corner-tl"></div>
              <div class="card-corner corner-tr"></div>
              <div class="card-corner corner-bl"></div>
              <div class="card-corner corner-br"></div>
              <div class="card-scan"></div>
              <div class="card-content">
                <div class="card-header">
                  <img :src="item.logo || defaultLogo" alt="" class="card-logo" />
                  <div class="card-title-area">
                    <h3 class="card-title" v-html="highlightText(item.name)"></h3>
                    <span class="card-cat">{{ item.catename }}</span>
                  </div>
                  <span class="card-type type-link">快链</span>
                </div>
                <div class="card-footer">
                  <span class="card-url">{{ item.url }}</span>
                  <svg class="card-arrow" viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M5 12h14M12 5l7 7-7 7"/>
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="initial-state" v-else>
        <div class="initial-logo">
          <div class="cyber-logo large">
            <span class="logo-text">HIX</span><span class="logo-text-accent">NAV</span>
          </div>
          <p class="initial-tagline">智能导航搜索系统</p>
        </div>
        <div class="hot-searches" v-if="hotSearches.length > 0">
          <h4>热门搜索</h4>
          <div class="hot-tags">
            <span v-for="(tag, idx) in hotSearches" :key="idx" class="hot-tag" @click="searchHot(tag)">{{ tag }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { search } from '@/api/search'

export default {
  name: 'SearchResultsView',
  data() {
    return {
      keyword: '',
      navResults: [],
      linkResults: [],
      loading: false,
      searched: false,
      searchTime: 0,
      showSuggestions: false,
      suggestions: [],
      searchTimer: null,
      defaultLogo: require('@/assets/logo.png'),
      hotSearches: ['百度', 'GitHub', 'Google', 'Vue', 'Docker'],
    }
  },
  computed: {
    totalResults() {
      return this.navResults.length + this.linkResults.length
    }
  },
  methods: {
    particleStyle(n) {
      const size = Math.random() * 4 + 2
      const left = Math.random() * 100
      const delay = Math.random() * 10
      const duration = Math.random() * 10 + 10
      return { width: size + 'px', height: size + 'px', left: left + '%', animationDelay: delay + 's', animationDuration: duration + 's' }
    },
    highlightText(text) {
      if (!text || !this.keyword) return text || ''
      const regex = new RegExp(`(${this.keyword.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')})`, 'gi')
      return text.replace(regex, '<mark class="highlight">$1</mark>')
    },
    onInput() {
      clearTimeout(this.searchTimer)
      if (this.keyword.trim().length >= 1) {
        this.searchTimer = setTimeout(() => { this.fetchSuggestions() }, 300)
      } else {
        this.suggestions = []
        this.showSuggestions = false
      }
    },
    async fetchSuggestions() {
      if (!this.keyword.trim()) return
      try {
        const res = await search({ keyword: this.keyword.trim() })
        if (res && res.data) {
          const all = [...(res.data.navs || []), ...(res.data.links || [])]
          this.suggestions = all.slice(0, 8)
          this.showSuggestions = this.suggestions.length > 0
        }
      } catch (e) {
        console.error('Suggestion fetch error:', e)
      }
    },
    selectSuggestion(item) {
      this.keyword = item.name
      this.showSuggestions = false
      this.doSearch()
    },
    async doSearch() {
      if (!this.keyword.trim()) return
      this.showSuggestions = false
      this.loading = true
      this.searched = true
      const startTime = Date.now()
      try {
        console.log('Searching for:', this.keyword.trim())
        const res = await search({ keyword: this.keyword.trim() })
        console.log('Search response:', res)
        this.searchTime = ((Date.now() - startTime) / 1000).toFixed(3)
        if (res && res.data) {
          console.log('Search data:', res.data)
          this.navResults = res.data.navs || []
          this.linkResults = res.data.links || []
          console.log('navResults:', this.navResults)
          console.log('linkResults:', this.linkResults)
        }
      } catch (e) {
        console.error('Search error:', e)
        this.$message.error('搜索失败，请稍后重试')
      } finally {
        this.loading = false
      }
    },
    jumpTo(url) {
      if (url) window.open(url, '_blank')
    },
    searchHot(tag) {
      this.keyword = tag
      this.doSearch()
    }
  },
  created() {
    const kw = this.$route.query.q
    if (kw) { this.keyword = kw; this.doSearch() }
  },
  mounted() {
    document.addEventListener('click', (e) => {
      if (!this.$refs.searchInput || !this.$refs.searchInput.closest('.search-box').contains(e.target)) {
        this.showSuggestions = false
      }
    })
  }
}
</script>

<style scoped>
.search-page { min-height: 100vh; background: #0a0e17; color: #e0e6f0; font-family: 'Segoe UI', 'PingFang SC', sans-serif; position: relative; overflow-x: hidden; }
.cyber-bg { position: fixed; top: 0; left: 0; width: 100%; height: 100%; z-index: 0; pointer-events: none; }
.grid-overlay { position: absolute; top: 0; left: 0; width: 100%; height: 100%; background-image: linear-gradient(rgba(0, 195, 255, 0.03) 1px, transparent 1px), linear-gradient(90deg, rgba(0, 195, 255, 0.03) 1px, transparent 1px); background-size: 50px 50px; animation: gridMove 20s linear infinite; }
@keyframes gridMove { 0% { transform: translate(0, 0); } 100% { transform: translate(50px, 50px); } }
.scan-line { position: absolute; top: 0; left: 0; width: 100%; height: 2px; background: linear-gradient(90deg, transparent, rgba(0, 195, 255, 0.4), transparent); animation: scanLine 4s linear infinite; }
@keyframes scanLine { 0% { top: 0; } 100% { top: 100%; } }
.floating-particles { position: absolute; width: 100%; height: 100%; }
.particle { position: absolute; bottom: -10px; background: rgba(0, 195, 255, 0.6); border-radius: 50%; animation: floatUp linear infinite; box-shadow: 0 0 6px rgba(0, 195, 255, 0.4); }
@keyframes floatUp { 0% { transform: translateY(0) scale(1); opacity: 0; } 10% { opacity: 1; } 90% { opacity: 1; } 100% { transform: translateY(-100vh) scale(0.5); opacity: 0; } }

.search-header { position: sticky; top: 0; z-index: 100; display: flex; align-items: center; padding: 16px 30px; background: rgba(10, 14, 23, 0.9); backdrop-filter: blur(20px); border-bottom: 1px solid rgba(0, 195, 255, 0.15); }
.logo-area { margin-right: 30px; flex-shrink: 0; }
.logo-link { text-decoration: none; }
.cyber-logo { display: flex; align-items: baseline; font-size: 24px; font-weight: 800; letter-spacing: 2px; }
.cyber-logo.large { font-size: 56px; justify-content: center; }
.logo-text { color: #00c3ff; text-shadow: 0 0 20px rgba(0, 195, 255, 0.5); }
.logo-text-accent { color: #fff; margin-left: 2px; }
.search-box { flex: 1; max-width: 700px; position: relative; }
.search-input-wrapper { display: flex; align-items: center; position: relative; border: 1px solid rgba(0, 195, 255, 0.3); border-radius: 8px; overflow: hidden; background: rgba(15, 23, 42, 0.8); transition: all 0.3s ease; }
.search-input-wrapper:focus-within { border-color: #00c3ff; box-shadow: 0 0 20px rgba(0, 195, 255, 0.2), inset 0 0 20px rgba(0, 195, 255, 0.05); }
.input-glow { position: absolute; top: 0; left: 0; width: 100%; height: 100%; background: linear-gradient(90deg, transparent, rgba(0, 195, 255, 0.1), transparent); animation: inputGlow 3s ease-in-out infinite; pointer-events: none; }
@keyframes inputGlow { 0%, 100% { opacity: 0; } 50% { opacity: 1; } }
.cyber-input { flex: 1; padding: 14px 20px; background: transparent; border: none; outline: none; color: #e0e6f0; font-size: 15px; letter-spacing: 0.5px; }
.cyber-input::placeholder { color: rgba(224, 230, 240, 0.4); }
.cyber-btn { display: flex; align-items: center; gap: 6px; padding: 14px 24px; background: linear-gradient(135deg, #00c3ff, #0077ff); border: none; color: #fff; font-size: 14px; font-weight: 600; cursor: pointer; transition: all 0.3s ease; letter-spacing: 1px; }
.cyber-btn:hover { background: linear-gradient(135deg, #00d4ff, #0088ff); box-shadow: 0 0 20px rgba(0, 195, 255, 0.4); }

.search-suggestions { position: absolute; top: calc(100% + 8px); left: 0; right: 0; background: rgba(15, 23, 42, 0.95); backdrop-filter: blur(20px); border: 1px solid rgba(0, 195, 255, 0.2); border-radius: 8px; overflow: hidden; z-index: 200; box-shadow: 0 10px 40px rgba(0, 0, 0, 0.5); }
.suggestion-item { display: flex; align-items: center; padding: 12px 16px; cursor: pointer; transition: background 0.2s; border-bottom: 1px solid rgba(255, 255, 255, 0.05); }
.suggestion-item:hover { background: rgba(0, 195, 255, 0.1); }
.sug-logo { width: 28px; height: 28px; border-radius: 6px; margin-right: 12px; object-fit: cover; }
.sug-info { flex: 1; display: flex; flex-direction: column; }
.sug-name { font-size: 14px; color: #e0e6f0; }
.sug-cat { font-size: 11px; color: rgba(224, 230, 240, 0.5); margin-top: 2px; }
.sug-type { font-size: 10px; padding: 2px 8px; border-radius: 4px; font-weight: 600; }
.type-nav { background: rgba(0, 195, 255, 0.15); color: #00c3ff; }
.type-link { background: rgba(139, 92, 246, 0.15); color: #8b5cf6; }

.search-results-container { position: relative; z-index: 1; max-width: 1200px; margin: 0 auto; padding: 30px; }
.loading-state { display: flex; justify-content: center; align-items: center; min-height: 400px; }
.cyber-loader { position: relative; width: 100px; height: 100px; display: flex; flex-direction: column; align-items: center; justify-content: center; }
.loader-ring { position: absolute; border: 2px solid transparent; border-radius: 50%; animation: loaderSpin linear infinite; }
.loader-ring:nth-child(1) { width: 80px; height: 80px; border-top-color: #00c3ff; animation-duration: 1.5s; }
.loader-ring:nth-child(2) { width: 60px; height: 60px; border-right-color: #8b5cf6; animation-duration: 1s; animation-direction: reverse; }
.loader-ring:nth-child(3) { width: 40px; height: 40px; border-bottom-color: #00ff88; animation-duration: 0.75s; }
@keyframes loaderSpin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
.loader-text { margin-top: 100px; font-size: 12px; letter-spacing: 4px; color: #00c3ff; animation: pulse 1.5s ease-in-out infinite; }
@keyframes pulse { 0%, 100% { opacity: 0.4; } 50% { opacity: 1; } }

.empty-state { text-align: center; padding: 80px 20px; }
.empty-icon { color: rgba(0, 195, 255, 0.3); margin-bottom: 20px; }
.empty-text { font-size: 18px; color: rgba(224, 230, 240, 0.6); }
.empty-text span { color: #00c3ff; }
.empty-hint { font-size: 14px; color: rgba(224, 230, 240, 0.3); margin-top: 10px; }

.initial-state { text-align: center; padding: 100px 20px; }
.initial-tagline { font-size: 16px; color: rgba(224, 230, 240, 0.4); letter-spacing: 6px; margin-top: 20px; }
.hot-searches { margin-top: 60px; }
.hot-searches h4 { font-size: 14px; color: rgba(224, 230, 240, 0.4); letter-spacing: 4px; margin-bottom: 20px; }
.hot-tags { display: flex; justify-content: center; flex-wrap: wrap; gap: 12px; }
.hot-tag { padding: 8px 20px; background: rgba(0, 195, 255, 0.08); border: 1px solid rgba(0, 195, 255, 0.2); border-radius: 20px; color: #00c3ff; font-size: 13px; cursor: pointer; transition: all 0.3s ease; }
.hot-tag:hover { background: rgba(0, 195, 255, 0.15); border-color: #00c3ff; box-shadow: 0 0 15px rgba(0, 195, 255, 0.2); }

.results-stats { margin-bottom: 30px; font-size: 13px; color: rgba(224, 230, 240, 0.5); }
.results-stats em { color: #00c3ff; font-style: normal; font-weight: 600; }
.stat-time { margin-left: 10px; }
.result-section { margin-bottom: 50px; }
.section-header { display: flex; align-items: center; margin-bottom: 24px; gap: 16px; }
.header-line { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(0, 195, 255, 0.3), transparent); }
.section-title { display: flex; align-items: center; gap: 10px; font-size: 18px; font-weight: 600; color: #00c3ff; letter-spacing: 2px; margin: 0; }
.count-badge { display: inline-flex; align-items: center; justify-content: center; min-width: 24px; height: 24px; padding: 0 8px; background: rgba(0, 195, 255, 0.15); border: 1px solid rgba(0, 195, 255, 0.3); border-radius: 12px; font-size: 12px; font-weight: 600; color: #00c3ff; }
.result-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(340px, 1fr)); gap: 20px; }

.cyber-card { position: relative; background: rgba(15, 23, 42, 0.6); border: 1px solid rgba(0, 195, 255, 0.15); border-radius: 12px; overflow: hidden; cursor: pointer; transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1); }
.cyber-card:hover { border-color: rgba(0, 195, 255, 0.5); transform: translateY(-4px); box-shadow: 0 10px 40px rgba(0, 0, 0, 0.3), 0 0 30px rgba(0, 195, 255, 0.1), inset 0 1px 0 rgba(0, 195, 255, 0.1); }
.card-corner { position: absolute; width: 16px; height: 16px; z-index: 2; }
.corner-tl { top: 0; left: 0; border-top: 2px solid #00c3ff; border-left: 2px solid #00c3ff; }
.corner-tr { top: 0; right: 0; border-top: 2px solid #00c3ff; border-right: 2px solid #00c3ff; }
.corner-bl { bottom: 0; left: 0; border-bottom: 2px solid #00c3ff; border-left: 2px solid #00c3ff; }
.corner-br { bottom: 0; right: 0; border-bottom: 2px solid #00c3ff; border-right: 2px solid #00c3ff; }
.card-scan { position: absolute; top: 0; left: 0; width: 100%; height: 2px; background: linear-gradient(90deg, transparent, rgba(0, 195, 255, 0.3), transparent); opacity: 0; transition: opacity 0.3s; }
.cyber-card:hover .card-scan { opacity: 1; animation: cardScan 2s linear infinite; }
@keyframes cardScan { 0% { top: 0; } 100% { top: 100%; } }
.card-content { padding: 20px; }
.card-header { display: flex; align-items: center; margin-bottom: 12px; }
.card-logo { width: 44px; height: 44px; border-radius: 10px; object-fit: cover; margin-right: 14px; border: 1px solid rgba(0, 195, 255, 0.2); background: rgba(0, 195, 255, 0.05); }
.card-title-area { flex: 1; min-width: 0; }
.card-title { font-size: 16px; font-weight: 600; color: #e0e6f0; margin: 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.card-cat { font-size: 11px; color: rgba(224, 230, 240, 0.4); margin-top: 4px; display: inline-block; }
.card-type { font-size: 10px; padding: 3px 10px; border-radius: 4px; font-weight: 600; letter-spacing: 1px; flex-shrink: 0; margin-left: 10px; }
.type-nav { background: rgba(0, 195, 255, 0.12); color: #00c3ff; border: 1px solid rgba(0, 195, 255, 0.2); }
.type-link { background: rgba(139, 92, 246, 0.12); color: #8b5cf6; border: 1px solid rgba(139, 92, 246, 0.2); }
.card-desc { font-size: 13px; color: rgba(224, 230, 240, 0.5); line-height: 1.6; margin: 0 0 14px 0; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.card-footer { display: flex; align-items: center; justify-content: space-between; padding-top: 12px; border-top: 1px solid rgba(255, 255, 255, 0.05); }
.card-url { font-size: 12px; color: rgba(0, 195, 255, 0.6); white-space: nowrap; overflow: hidden; text-overflow: ellipsis; flex: 1; margin-right: 10px; }
.card-arrow { color: rgba(0, 195, 255, 0.4); transition: all 0.3s; flex-shrink: 0; }
.cyber-card:hover .card-arrow { color: #00c3ff; transform: translateX(4px); }
.link-card .card-type { background: rgba(139, 92, 246, 0.12); color: #8b5cf6; border-color: rgba(139, 92, 246, 0.2); }
:deep(.highlight) { background: rgba(0, 195, 255, 0.2); color: #00e5ff; padding: 1px 4px; border-radius: 3px; font-weight: 600; }

@media (max-width: 768px) {
  .search-header { flex-direction: column; padding: 12px 16px; gap: 12px; }
  .logo-area { margin-right: 0; }
  .search-box { max-width: 100%; }
  .search-results-container { padding: 20px 16px; }
  .result-grid { grid-template-columns: 1fr; }
  .cyber-logo.large { font-size: 36px; }
}
</style>
