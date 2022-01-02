package com.EquityTrackr.backend.investor;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.concurrent.ExecutionException;

@RestController
@RequestMapping(path = "api/v1/investor")
public class InvestorController {

    private final InvestorService investorService;

    @Autowired
    public InvestorController(InvestorService investorService) {
        this.investorService = investorService;
    }


    @GetMapping
    public Investor getInvestor(@RequestParam String name) throws InterruptedException, ExecutionException {
        return investorService.getInvestor(name);
    }

    @PostMapping
    @ResponseBody
    public Investor createInvestor(@RequestBody Investor investor) throws InterruptedException, ExecutionException {
        return investorService.saveInvestor(investor);
    }

//    @PutMapping(path = "{investorId}")
//    public String updateInvestor(@PathVariable("studentId") Long studentId, @RequestBody Investor investor) throws InterruptedException, ExecutionException {
//        return investorService.updateInvestor(investor);
//    }

    @DeleteMapping
    public String deleteInvestor(@RequestParam String name) {
        return investorService.deleteInvestor(name);
    }
}