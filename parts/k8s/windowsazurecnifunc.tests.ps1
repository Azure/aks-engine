. $PSScriptRoot\windowsazurecnifunc.ps1

Describe 'GetBroadestRangesForEachAddress' {

    It "Values '<Values>' should return '<Expected>'" -TestCases @(
        @{ Values = @('10.240.0.0/12', '10.0.0.0/8'); Expected = @('10.0.0.0/8', '10.240.0.0/12')}
        @{ Values = @('10.0.0.0/8', '10.0.0.0/16'); Expected = @('10.0.0.0/8')}
        @{ Values = @('10.0.0.0/16', '10.240.0.0/12', '10.0.0.0/8' ); Expected = @('10.0.0.0/8', '10.240.0.0/12')}
        @{ Values = @(); Expected = @()}
        @{ Values = @('foobar'); Expected = @()}
    ){
        param ($Values, $Expected)

        $actual = GetBroadestRangesForEachAddress -values $Values
        $actual | Should -Be $Expected
    }
}