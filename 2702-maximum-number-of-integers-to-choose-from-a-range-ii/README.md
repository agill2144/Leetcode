<h2><a href="https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-ii">Maximum Number of Integers to Choose From a Range II</a></h2> <img src='https://img.shields.io/badge/Difficulty-Medium-orange' alt='Difficulty: Medium' /><hr><p>You are given an integer array <code>banned</code> and two integers <code>n</code> and <code>maxSum</code>. You are choosing some number of integers following the below rules:</p>

<ul>
	<li>The chosen integers have to be in the range <code>[1, n]</code>.</li>
	<li>Each integer can be chosen <strong>at most once</strong>.</li>
	<li>The chosen integers should not be in the array <code>banned</code>.</li>
	<li>The sum of the chosen integers should not exceed <code>maxSum</code>.</li>
</ul>

<p>Return <em>the <strong>maximum</strong> number of integers you can choose following the mentioned rules</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> banned = [1,4,6], n = 6, maxSum = 4
<strong>Output:</strong> 1
<strong>Explanation:</strong> You can choose the integer 3.
3 is in the range [1, 6], and do not appear in banned. The sum of the chosen integers is 3, which does not exceed maxSum.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> banned = [4,3,5,6], n = 7, maxSum = 18
<strong>Output:</strong> 3
<strong>Explanation:</strong> You can choose the integers 1, 2, and 7.
All these integers are in the range [1, 7], all do not appear in banned, and their sum is 18, which does not exceed maxSum.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= banned.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= banned[i] &lt;= n &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= maxSum &lt;= 10<sup>15</sup></code></li>
</ul>
